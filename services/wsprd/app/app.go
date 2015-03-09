// The app package contains the struct that keeps per javascript app state and handles translating
// javascript requests to veyron requests and vice versa.
package app

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"reflect"
	"sync"
	"time"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/ipc"
	"v.io/v23/naming"
	"v.io/v23/options"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/vdlroot/signature"
	"v.io/v23/verror"
	"v.io/v23/vom"
	"v.io/v23/vtrace"
	"v.io/x/lib/vlog"
	vsecurity "v.io/x/ref/security"
	"v.io/x/ref/services/wsprd/ipc/server"
	"v.io/x/ref/services/wsprd/lib"
	"v.io/x/ref/services/wsprd/namespace"
	"v.io/x/ref/services/wsprd/principal"
)

// pkgPath is the prefix os errors in this package.
const pkgPath = "v.io/x/ref/services/wsprd/app"

// Errors
var (
	marshallingError       = verror.Register(pkgPath+".marshallingError", verror.NoRetry, "{1} {2} marshalling error {_}")
	noResults              = verror.Register(pkgPath+".noResults", verror.NoRetry, "{1} {2} no results from call {_}")
	badCaveatType          = verror.Register(pkgPath+".badCaveatType", verror.NoRetry, "{1} {2} bad caveat type {_}")
	unknownBlessings       = verror.Register(pkgPath+".unknownBlessings", verror.NoRetry, "{1} {2} unknown public id {_}")
	invalidBlessingsHandle = verror.Register(pkgPath+".invalidBlessingsHandle", verror.NoRetry, "{1} {2} invalid blessings handle {_}")
)

// TODO(bjornick,nlacasse): Remove the retryTimeout flag once we able
// to pass it in from javascript. For now all RPCs have the same
// retryTimeout, set by command line flag.
var retryTimeout *int

func init() {
	// TODO(bjornick,nlacasse): Remove the retryTimeout flag once we able
	// to pass it in from javascript. For now all RPCs have the same
	// retryTimeout, set by command line flag.
	retryTimeout = flag.Int("retry-timeout", 2, "Duration in seconds to retry starting an RPC call. 0 means never retry.")
}

type outstandingRequest struct {
	stream *outstandingStream
	cancel context.CancelFunc
}

// Controller represents all the state of a Veyron Web App.  This is the struct
// that is in charge performing all the veyron options.
type Controller struct {
	// Protects everything.
	// TODO(bjornick): We need to split this up.
	sync.Mutex

	// The context of this controller.
	ctx *context.T

	// The cleanup function for this controller.
	cancel context.CancelFunc

	// The ipc.ListenSpec to use with server.Listen
	listenSpec *ipc.ListenSpec

	// Used to generate unique ids for requests initiated by the proxy.
	// These ids will be even so they don't collide with the ids generated
	// by the client.
	lastGeneratedId int32

	// Used to keep track of data (streams and cancellation functions) for
	// outstanding requests.
	outstandingRequests map[int32]*outstandingRequest

	// Maps flowids to the server that owns them.
	flowMap map[int32]*server.Server

	// A manager that Handles fetching and caching signature of remote services
	signatureManager lib.SignatureManager

	// We maintain multiple Veyron server per pipe for serving JavaScript
	// services.
	servers map[uint32]*server.Server

	// Creates a client writer for a given flow.  This is a member so that tests can override
	// the default implementation.
	writerCreator func(id int32) lib.ClientWriter

	// Store for all the Blessings that javascript has a handle to.
	blessingsStore *principal.JSBlessingsHandles

	// reservedServices contains a map of reserved service names.  These
	// are objects that serve requests in wspr without actually making
	// an outgoing rpc call.
	reservedServices map[string]ipc.Invoker
}

// NewController creates a new Controller.  writerCreator will be used to create a new flow for rpcs to
// javascript server.
func NewController(ctx *context.T, writerCreator func(id int32) lib.ClientWriter, listenSpec *ipc.ListenSpec, namespaceRoots []string, p security.Principal) (*Controller, error) {
	ctx, cancel := context.WithCancel(ctx)

	if namespaceRoots != nil {
		var err error
		ctx, _, err = v23.SetNewNamespace(ctx, namespaceRoots...)
		if err != nil {
			return nil, err
		}
	}

	ctx, _ = vtrace.SetNewTrace(ctx)

	ctx, err := v23.SetPrincipal(ctx, p)
	if err != nil {
		return nil, err
	}

	controller := &Controller{
		ctx:            ctx,
		cancel:         cancel,
		writerCreator:  writerCreator,
		listenSpec:     listenSpec,
		blessingsStore: principal.NewJSBlessingsHandles(),
	}

	controllerInvoker, err := ipc.ReflectInvoker(ControllerServer(controller))
	if err != nil {
		return nil, err
	}
	namespaceInvoker, err := ipc.ReflectInvoker(namespace.New(ctx))
	if err != nil {
		return nil, err
	}
	controller.reservedServices = map[string]ipc.Invoker{
		"__controller": controllerInvoker,
		"__namespace":  namespaceInvoker,
	}

	controller.setup()
	return controller, nil
}

// finishCall waits for the call to finish and write out the response to w.
func (c *Controller) finishCall(ctx *context.T, w lib.ClientWriter, clientCall ipc.ClientCall, msg *VeyronRPCRequest, span vtrace.Span) {
	if msg.IsStreaming {
		for {
			var item interface{}
			if err := clientCall.Recv(&item); err != nil {
				if err == io.EOF {
					break
				}
				w.Error(err) // Send streaming error as is
				return
			}
			vomItem, err := lib.VomEncode(item)
			if err != nil {
				w.Error(verror.New(marshallingError, ctx, item, err))
				continue
			}
			if err := w.Send(lib.ResponseStream, vomItem); err != nil {
				w.Error(verror.New(marshallingError, ctx, item))
			}
		}
		if err := w.Send(lib.ResponseStreamClose, nil); err != nil {
			w.Error(verror.New(marshallingError, ctx, "ResponseStreamClose"))
		}
	}
	results := make([]*vdl.Value, msg.NumOutArgs)
	// This array will have pointers to the values in results.
	resultptrs := make([]interface{}, msg.NumOutArgs)
	for i := range results {
		resultptrs[i] = &results[i]
	}
	if err := clientCall.Finish(resultptrs...); err != nil {
		// return the call system error as is
		w.Error(err)
		return
	}
	c.sendRPCResponse(ctx, w, span, results)
}

func (c *Controller) sendRPCResponse(ctx *context.T, w lib.ClientWriter, span vtrace.Span, results []*vdl.Value) {
	span.Finish()
	traceRecord := vtrace.GetStore(ctx).TraceRecord(span.Trace())

	response := VeyronRPCResponse{
		OutArgs: results,
		TraceResponse: vtrace.Response{
			Trace: *traceRecord,
		},
	}
	encoded, err := lib.VomEncode(response)
	if err != nil {
		w.Error(err)
		return
	}
	if err := w.Send(lib.ResponseFinal, encoded); err != nil {
		w.Error(verror.Convert(marshallingError, ctx, err))
	}
}

func (c *Controller) startCall(ctx *context.T, w lib.ClientWriter, msg *VeyronRPCRequest, inArgs []interface{}) (ipc.ClientCall, error) {
	methodName := lib.UppercaseFirstCharacter(msg.Method)
	retryTimeoutOpt := options.RetryTimeout(time.Duration(*retryTimeout) * time.Second)
	clientCall, err := v23.GetClient(ctx).StartCall(ctx, msg.Name, methodName, inArgs, retryTimeoutOpt)
	if err != nil {
		return nil, fmt.Errorf("error starting call (name: %v, method: %v, args: %v): %v", msg.Name, methodName, inArgs, err)
	}

	return clientCall, nil
}

// Implements the serverHelper interface

// CreateNewFlow creats a new server flow that will be used to write out
// streaming messages to Javascript.
func (c *Controller) CreateNewFlow(s *server.Server, stream ipc.Stream) *server.Flow {
	c.Lock()
	defer c.Unlock()
	id := c.lastGeneratedId
	c.lastGeneratedId += 2
	c.flowMap[id] = s
	os := newStream()
	os.init(stream)
	c.outstandingRequests[id] = &outstandingRequest{
		stream: os,
	}
	return &server.Flow{ID: id, Writer: c.writerCreator(id)}
}

// CleanupFlow removes the bookkeping for a previously created flow.
func (c *Controller) CleanupFlow(id int32) {
	c.Lock()
	request := c.outstandingRequests[id]
	delete(c.outstandingRequests, id)
	delete(c.flowMap, id)
	c.Unlock()
	if request != nil && request.stream != nil {
		request.stream.end()
		request.stream.waitUntilDone()
	}
}

// RT returns the runtime of the app.
func (c *Controller) Context() *context.T {
	return c.ctx
}

// AddBlessings adds the Blessings to the local blessings store and returns
// the handle to it.  This function exists because JS only has
// a handle to the blessings to avoid shipping the certificate forest
// to JS and back.
func (c *Controller) AddBlessings(blessings security.Blessings) int32 {
	return c.blessingsStore.Add(blessings)
}

// Cleanup cleans up any outstanding rpcs.
func (c *Controller) Cleanup() {
	vlog.VI(0).Info("Cleaning up controller")
	c.Lock()

	for _, request := range c.outstandingRequests {
		if request.cancel != nil {
			request.cancel()
		}
		if request.stream != nil {
			request.stream.end()
		}
	}

	servers := []*server.Server{}
	for _, server := range c.servers {
		servers = append(servers, server)
	}

	c.Unlock()

	// We must unlock before calling server.Stop otherwise it can deadlock.
	for _, server := range servers {
		server.Stop()
	}

	c.cancel()
}

func (c *Controller) setup() {
	c.signatureManager = lib.NewSignatureManager()
	c.outstandingRequests = make(map[int32]*outstandingRequest)
	c.flowMap = make(map[int32]*server.Server)
	c.servers = make(map[uint32]*server.Server)
}

// SendOnStream writes data on id's stream.  The actual network write will be
// done asynchronously.  If there is an error, it will be sent to w.
func (c *Controller) SendOnStream(id int32, data string, w lib.ClientWriter) {
	c.Lock()
	request := c.outstandingRequests[id]
	if request == nil || request.stream == nil {
		vlog.Errorf("unknown stream: %d", id)
		c.Unlock()
		return
	}
	stream := request.stream
	c.Unlock()
	stream.send(data, w)
}

// SendVeyronRequest makes a veyron request for the given flowId.  If signal is non-nil, it will receive
// the call object after it has been constructed.
func (c *Controller) sendVeyronRequest(ctx *context.T, id int32, msg *VeyronRPCRequest, inArgs []interface{}, w lib.ClientWriter, stream *outstandingStream, span vtrace.Span) {
	sig, err := c.getSignature(ctx, msg.Name)
	if err != nil {
		w.Error(err)
		return
	}
	methName := lib.UppercaseFirstCharacter(msg.Method)
	methSig, ok := signature.FirstMethod(sig, methName)
	if !ok {
		w.Error(fmt.Errorf("method %q not found in signature: %#v", methName, sig))
		return
	}
	if len(methSig.InArgs) != len(inArgs) {
		w.Error(fmt.Errorf("invalid number of arguments, expected: %v, got:%v", methSig, *msg))
		return
	}

	// We have to make the start call synchronous so we can make sure that we populate
	// the call map before we can Handle a recieve call.
	call, err := c.startCall(ctx, w, msg, inArgs)
	if err != nil {
		w.Error(verror.Convert(verror.ErrInternal, ctx, err))
		return
	}

	if stream != nil {
		stream.init(call)
	}

	c.finishCall(ctx, w, call, msg, span)
	c.Lock()
	if request, ok := c.outstandingRequests[id]; ok {
		delete(c.outstandingRequests, id)
		if request.cancel != nil {
			request.cancel()
		}
	}
	c.Unlock()
}

// TODO(mattr): This is a very limited implementation of ServerCall,
// but currently none of the methods the controller exports require
// any of this context information.
type localCall struct {
	ctx  *context.T
	vrpc *VeyronRPCRequest
	tags []*vdl.Value
	w    lib.ClientWriter
}

func (l *localCall) Send(item interface{}) error {
	vomItem, err := lib.VomEncode(item)
	if err != nil {
		err = verror.New(marshallingError, l.ctx, item, err)
		l.w.Error(err)
		return err
	}
	if err := l.w.Send(lib.ResponseStream, vomItem); err != nil {
		err = verror.New(marshallingError, l.ctx, item)
		l.w.Error(err)
		return err
	}
	return nil
}
func (l *localCall) Recv(interface{}) error                          { return nil }
func (l *localCall) GrantedBlessings() security.Blessings            { return security.Blessings{} }
func (l *localCall) Server() ipc.Server                              { return nil }
func (l *localCall) Context() *context.T                             { return l.ctx }
func (l *localCall) Timestamp() (t time.Time)                        { return }
func (l *localCall) Method() string                                  { return l.vrpc.Method }
func (l *localCall) MethodTags() []*vdl.Value                        { return l.tags }
func (l *localCall) Name() string                                    { return l.vrpc.Name }
func (l *localCall) Suffix() string                                  { return "" }
func (l *localCall) LocalDischarges() map[string]security.Discharge  { return nil }
func (l *localCall) RemoteDischarges() map[string]security.Discharge { return nil }
func (l *localCall) LocalPrincipal() security.Principal              { return nil }
func (l *localCall) LocalBlessings() security.Blessings              { return security.Blessings{} }
func (l *localCall) RemoteBlessings() security.Blessings             { return security.Blessings{} }
func (l *localCall) LocalEndpoint() naming.Endpoint                  { return nil }
func (l *localCall) RemoteEndpoint() naming.Endpoint                 { return nil }
func (l *localCall) VanadiumContext() *context.T                     { return l.ctx }

func (c *Controller) handleInternalCall(ctx *context.T, invoker ipc.Invoker, msg *VeyronRPCRequest, decoder *vom.Decoder, w lib.ClientWriter, span vtrace.Span) {
	argptrs, tags, err := invoker.Prepare(msg.Method, int(msg.NumInArgs))
	if err != nil {
		w.Error(verror.Convert(verror.ErrInternal, ctx, err))
		return
	}
	for _, argptr := range argptrs {
		if err := decoder.Decode(argptr); err != nil {
			w.Error(verror.Convert(verror.ErrInternal, ctx, err))
			return
		}
	}
	results, err := invoker.Invoke(msg.Method, &localCall{ctx, msg, tags, w}, argptrs)
	if err != nil {
		w.Error(verror.Convert(verror.ErrInternal, ctx, err))
		return
	}
	if msg.IsStreaming {
		if err := w.Send(lib.ResponseStreamClose, nil); err != nil {
			w.Error(verror.New(marshallingError, ctx, "ResponseStreamClose"))
		}
	}

	// Convert results from []interface{} to []*vdl.Value.
	vresults := make([]*vdl.Value, len(results))
	for i, res := range results {
		vv, err := vdl.ValueFromReflect(reflect.ValueOf(res))
		if err != nil {
			w.Error(verror.Convert(verror.ErrInternal, ctx, err))
			return
		}
		vresults[i] = vv
	}
	c.sendRPCResponse(ctx, w, span, vresults)
}

// HandleCaveatValidationResponse handles the response to caveat validation
// requests.
func (c *Controller) HandleCaveatValidationResponse(id int32, data string) {
	c.Lock()
	server := c.flowMap[id]
	c.Unlock()
	if server == nil {
		vlog.Errorf("unexpected result from JavaScript. No server found matching id %d.", id)
		return // ignore unknown server
	}
	server.HandleCaveatValidationResponse(id, data)
}

// HandleVeyronRequest starts a veyron rpc and returns before the rpc has been completed.
func (c *Controller) HandleVeyronRequest(ctx *context.T, id int32, data string, w lib.ClientWriter) {
	binbytes, err := hex.DecodeString(data)
	if err != nil {
		w.Error(verror.Convert(verror.ErrInternal, ctx, err))
		return
	}
	decoder, err := vom.NewDecoder(bytes.NewReader(binbytes))
	if err != nil {
		w.Error(verror.Convert(verror.ErrInternal, ctx, err))
		return
	}

	var msg VeyronRPCRequest
	if err := decoder.Decode(&msg); err != nil {
		w.Error(verror.Convert(verror.ErrInternal, ctx, err))
		return
	}
	vlog.VI(2).Infof("VeyronRPC: %s.%s(..., streaming=%v)", msg.Name, msg.Method, msg.IsStreaming)
	spanName := fmt.Sprintf("<wspr>%q.%s", msg.Name, msg.Method)
	ctx, span := vtrace.SetContinuedTrace(ctx, spanName, msg.TraceRequest)

	var cctx *context.T
	var cancel context.CancelFunc

	// TODO(mattr): To be consistent with go, we should not ignore 0 timeouts.
	// However as a rollout strategy we must, otherwise there is a circular
	// dependency between the WSPR change and the JS change that will follow.
	if msg.Deadline.IsZero() {
		cctx, cancel = context.WithCancel(ctx)
	} else {
		cctx, cancel = context.WithDeadline(ctx, msg.Deadline.Time)
	}

	// If this message is for an internal service, do a short-circuit dispatch here.
	if invoker, ok := c.reservedServices[msg.Name]; ok {
		go c.handleInternalCall(ctx, invoker, &msg, decoder, w, span)
		return
	}

	inArgs := make([]interface{}, msg.NumInArgs)
	for i := range inArgs {
		var v *vdl.Value
		if err := decoder.Decode(&v); err != nil {
			w.Error(err)
			return
		}
		inArgs[i] = v
	}

	request := &outstandingRequest{
		cancel: cancel,
	}
	if msg.IsStreaming {
		// If this rpc is streaming, we would expect that the client would try to send
		// on this stream.  Since the initial handshake is done asynchronously, we have
		// to put the outstanding stream in the map before we make the async call so that
		// the future send know which queue to write to, even if the client call isn't
		// actually ready yet.
		request.stream = newStream()
	}
	c.Lock()
	c.outstandingRequests[id] = request
	go c.sendVeyronRequest(cctx, id, &msg, inArgs, w, request.stream, span)
	c.Unlock()
}

// HandleVeyronCancellation cancels the request corresponding to the
// given id if it is still outstanding.
func (c *Controller) HandleVeyronCancellation(id int32) {
	c.Lock()
	defer c.Unlock()
	if request, ok := c.outstandingRequests[id]; ok && request.cancel != nil {
		request.cancel()
	}
}

// CloseStream closes the stream for a given id.
func (c *Controller) CloseStream(id int32) {
	c.Lock()
	defer c.Unlock()
	if request, ok := c.outstandingRequests[id]; ok && request.stream != nil {
		request.stream.end()
		return
	}
	vlog.Errorf("close called on non-existent call: %v", id)
}

func (c *Controller) maybeCreateServer(serverId uint32) (*server.Server, error) {
	c.Lock()
	defer c.Unlock()
	if server, ok := c.servers[serverId]; ok {
		return server, nil
	}
	server, err := server.NewServer(serverId, c.listenSpec, c)
	if err != nil {
		return nil, err
	}
	c.servers[serverId] = server
	return server, nil
}

// HandleLookupResponse handles the result of a Dispatcher.Lookup call that was
// run by the Javascript server.
func (c *Controller) HandleLookupResponse(id int32, data string) {
	c.Lock()
	server := c.flowMap[id]
	c.Unlock()
	if server == nil {
		vlog.Errorf("unexpected result from JavaScript. No channel "+
			"for MessageId: %d exists. Ignoring the results.", id)
		//Ignore unknown responses that don't belong to any channel
		return
	}
	server.HandleLookupResponse(id, data)
}

// HandleAuthResponse handles the result of a Authorizer.Authorize call that was
// run by the Javascript server.
func (c *Controller) HandleAuthResponse(id int32, data string) {
	c.Lock()
	server := c.flowMap[id]
	c.Unlock()
	if server == nil {
		vlog.Errorf("unexpected result from JavaScript. No channel "+
			"for MessageId: %d exists. Ignoring the results.", id)
		//Ignore unknown responses that don't belong to any channel
		return
	}
	server.HandleAuthResponse(id, data)
}

// Serve instructs WSPR to start listening for calls on behalf
// of a javascript server.
func (c *Controller) Serve(_ ipc.ServerCall, name string, serverId uint32) error {
	server, err := c.maybeCreateServer(serverId)
	if err != nil {
		return verror.Convert(verror.ErrInternal, nil, err)
	}
	vlog.VI(2).Infof("serving under name: %q", name)
	if err := server.Serve(name); err != nil {
		return verror.Convert(verror.ErrInternal, nil, err)
	}
	return nil
}

// Stop instructs WSPR to stop listening for calls for the
// given javascript server.
func (c *Controller) Stop(_ ipc.ServerCall, serverId uint32) error {
	c.Lock()
	server := c.servers[serverId]
	if server == nil {
		c.Unlock()
		return nil
	}
	delete(c.servers, serverId)
	c.Unlock()

	server.Stop()
	return nil
}

// AddName adds a published name to an existing server.
func (c *Controller) AddName(_ ipc.ServerCall, serverId uint32, name string) error {
	// Create a server for the pipe, if it does not exist already
	server, err := c.maybeCreateServer(serverId)
	if err != nil {
		return verror.Convert(verror.ErrInternal, nil, err)
	}
	// Add name
	if err := server.AddName(name); err != nil {
		return verror.Convert(verror.ErrInternal, nil, err)
	}
	return nil
}

// RemoveName removes a published name from an existing server.
func (c *Controller) RemoveName(_ ipc.ServerCall, serverId uint32, name string) error {
	// Create a server for the pipe, if it does not exist already
	server, err := c.maybeCreateServer(serverId)
	if err != nil {
		return verror.Convert(verror.ErrInternal, nil, err)
	}
	// Remove name
	server.RemoveName(name)
	// Remove name from signature cache as well
	c.signatureManager.FlushCacheEntry(name)
	return nil
}

// HandleServerResponse handles the completion of outstanding calls to JavaScript services
// by filling the corresponding channel with the result from JavaScript.
func (c *Controller) HandleServerResponse(id int32, data string) {
	c.Lock()
	server := c.flowMap[id]
	c.Unlock()
	if server == nil {
		vlog.Errorf("unexpected result from JavaScript. No channel "+
			"for MessageId: %d exists. Ignoring the results.", id)
		//Ignore unknown responses that don't belong to any channel
		return
	}
	server.HandleServerResponse(id, data)
}

// parseVeyronRequest parses a json rpc request into a VeyronRPCRequest object.
func (c *Controller) parseVeyronRequest(data string) (*VeyronRPCRequest, error) {
	var msg VeyronRPCRequest
	if err := lib.VomDecode(data, &msg); err != nil {
		return nil, err
	}
	vlog.VI(2).Infof("VeyronRPCRequest: %s.%s(..., streaming=%v)", msg.Name, msg.Method, msg.IsStreaming)
	return &msg, nil
}

// getSignature uses the signature manager to get and cache the signature of a remote server.
func (c *Controller) getSignature(ctx *context.T, name string) ([]signature.Interface, error) {
	retryTimeoutOpt := options.RetryTimeout(time.Duration(*retryTimeout) * time.Second)
	return c.signatureManager.Signature(ctx, name, retryTimeoutOpt)
}

// Signature uses the signature manager to get and cache the signature of a remote server.
func (c *Controller) Signature(call ipc.ServerCall, name string) ([]signature.Interface, error) {
	return c.getSignature(call.Context(), name)
}

// UnlinkJSBlessings removes the given blessings from the blessings store.
func (c *Controller) UnlinkJSBlessings(_ ipc.ServerCall, handle int32) error {
	c.blessingsStore.Remove(handle)
	return nil
}

// BlessPublicKey creates a new blessing.
func (c *Controller) BlessPublicKey(_ ipc.ServerCall,
	handle int32,
	caveats []security.Caveat,
	duration time.Duration,
	extension string) (int32, string, error) {
	var blessee security.Blessings
	if blessee = c.blessingsStore.Get(handle); blessee.IsZero() {
		return 0, "", verror.New(invalidBlessingsHandle, nil)
	}

	expiryCav, err := security.ExpiryCaveat(time.Now().Add(duration))
	if err != nil {
		return 0, "", err
	}
	caveats = append(caveats, expiryCav)

	// TODO(ataly, ashankar, bjornick): Currently the Bless operation is carried
	// out using the Default blessing in this principal's blessings store. We
	// should change this so that the JS blessing request can also specify the
	// blessing to be used for the Bless operation.
	p := v23.GetPrincipal(c.ctx)
	key := blessee.PublicKey()
	blessing := p.BlessingStore().Default()
	blessings, err := p.Bless(key, blessing, extension, caveats[0], caveats[1:]...)
	if err != nil {
		return 0, "", err
	}
	handle = c.blessingsStore.Add(blessings)
	encodedKey, err := principal.EncodePublicKey(blessings.PublicKey())
	if err != nil {
		return 0, "", err
	}
	return handle, encodedKey, nil
}

// CreateBlessings creates a new principal self-blessed with the given extension.
func (c *Controller) CreateBlessings(_ ipc.ServerCall,
	extension string) (int32, string, error) {
	p, err := vsecurity.NewPrincipal()
	if err != nil {
		return 0, "", verror.Convert(verror.ErrInternal, nil, err)
	}
	blessings, err := p.BlessSelf(extension)
	if err != nil {
		return 0, "", verror.Convert(verror.ErrInternal, nil, err)
	}

	handle := c.blessingsStore.Add(blessings)
	encodedKey, err := principal.EncodePublicKey(blessings.PublicKey())
	if err != nil {
		return 0, "", err
	}
	return handle, encodedKey, nil
}

func (c *Controller) RemoteBlessings(call ipc.ServerCall, name, method string) ([]string, error) {
	vlog.VI(2).Infof("requesting remote blessings for %q", name)

	cctx, cancel := context.WithTimeout(call.Context(), 5*time.Second)
	defer cancel()

	clientCall, err := v23.GetClient(cctx).StartCall(cctx, name, method, nil)
	if err != nil {
		return nil, verror.Convert(verror.ErrInternal, cctx, err)
	}

	blessings, _ := clientCall.RemoteBlessings()
	return blessings, nil
}
