// This file was auto-generated by the veyron vdl tool.
// Source: syncgroup.vdl

// Package syncgroup provides the means for Stores to set up SyncGroups for
// subsequent synchronization of objects between them.
//
// The intent is that SyncGroup objects are created and administered by
// SyncGroup servers, even though they are subsequently mirrored among SyncGroup
// members by the normal synchronization mechanisms.
//
// SyncGroupServer also aids in discovering members of a particular SyncGroup.
// SyncGroupServer maintains the names of joiners for a SyncGroup that are in
// turn accessible to all members.  Each SyncGroup is also associated with a
// set of mount tables.  A Store that joins a SyncGroup must advertise its name
// (and optionally its SyncGroups) in these mount tables.  Stores are expected
// also optionally to advertise the SyncGroups they join in the local
// neighbourhood.
package syncgroup

import (
	"veyron2/security"

	"veyron2/services/security/access"

	"veyron2/storage"

	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_context "veyron2/context"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_vdlutil "veyron2/vdl/vdlutil"
	_gen_wiretype "veyron2/wiretype"
)

// An ID is a globally unique identifier for a SyncGroup.
type ID storage.ID

// A SyncGroupInfo is the conceptual state of a SyncGroup object.
type SyncGroupInfo struct {
	Name    string          // Global Veyron name of object.
	Config  SyncGroupConfig // Configuration parameters of this SyncGroup.
	RootOID storage.ID      // ID of object at root of SyncGroup's tree.
	ETag    string          // Version ID for concurrency control.
	// The SyncGroup's object ID, which is chosen by the creating SyncGroupServer
	// and is globally unique.
	SGOID ID
	// A map from joiner names to the associated metaData for devices that
	// have called Join() or Create() and not subsequently called Leave()
	// or had Eject() called on them.  The map returned by the calls below
	// may contain only a subset of joiners if the number is large.
	Joiners map[NameIdentity]JoinerMetaData
}

// A SyncGroupConfig contains some fields of SyncGroupInfo that
// are passed at create time, but which can be changed later.
type SyncGroupConfig struct {
	Desc         string                      // Human readable description.
	PathPatterns []string                    // Global path patterns.
	Options      map[string]_gen_vdlutil.Any // Options for future evolution.
	ACL          security.ACL                // The object's ACL.
	// Mount tables used to advertise for synchronization.
	// Typically, we will have only one entry.  However, an array allows
	// mount tables to be changed over time.
	MountTables []string
}

// A JoinerMetaData contains the non-name information stored per joiner.
type JoinerMetaData struct {
	// SyncPriority is a hint to bias the choice of syncing partners.
	// Members of the SyncGroup should choose to synchronize more often
	// with partners with lower values.
	SyncPriority int32
}

// A NameIdentity gives a Veyron name and identity for a joiner.
// TODO(m3b):  When names include an identity, this should become a single
// string.
type NameIdentity struct {
	Name     string // Global name of joiner.
	Identity string // Security identity of the joiner.
}

// TODO(bprosnitz) Remove this line once signatures are updated to use typevals.
// It corrects a bug where _gen_wiretype is unused in VDL pacakges where only bootstrap types are used on interfaces.
const _ = _gen_wiretype.TypeIDInvalid

// SyncGroupServer is the collection of calls on SyncGroup objects at
// a SyncGroup server.  The calls used most often, like Create and Join, are
// used almost exclusively by the Store.  Clients typically call the Store to
// cause these things to happen.
//
// Calls starting with "Set" take an eTag value that may be either empty, or
// the value of ETag from a recent response to Get(), Watch(), or GetACL().
// SyncGroupServer is the interface the client binds and uses.
// SyncGroupServer_ExcludingUniversal is the interface without internal framework-added methods
// to enable embedding without method collisions.  Not to be used directly by clients.
type SyncGroupServer_ExcludingUniversal interface {
	// Object provides access control for Veyron objects.
	access.Object_ExcludingUniversal
	// Create creates this SyncGroup with the given arguments, and if
	// joiner.Name!="", with {joiner, metaData} in its Joiners map.  It is
	// expected that acl will give Read and Write access to any device that
	// the administrator expects to join and sync; if the acl is empty, a
	// default ACL giving access only to the caller is used.  On success,
	// Create returns the SyncGroupInfo of the newly created object.
	//
	// Requires: this SyncGroup must not exist;
	// the caller must have write permission at the SyncGroup server;
	// the caller's identity must be a prefix of the SyncGroup's name.
	// Beware that for Create(), the access label is matched against a
	// server-wide ACL; for all other calls the access label is matched
	// against the object's ACL.
	Create(ctx _gen_context.T, createArgs SyncGroupConfig, rootOID storage.ID, joiner NameIdentity, metaData JoinerMetaData, opts ..._gen_ipc.CallOpt) (reply SyncGroupInfo, err error)
	// Join adds {joiner, metaData} to the SyncGroup's Joiners map and
	// returns the SyncGroupInfo for this SyncGroup.  The act of joining
	// allows other devices to find the caller, which is still required to
	// have read+write access on the SyncGroup to participate in
	// synchronization.  A device may call Join again with the same
	// NameIdentity in order to change metaData.
	// For SyncGroups with large numbers of joiners, Join may return
	// a subset of Joiners.
	//
	// Requires: this SyncGroup must exist;
	// the caller must have both read and write permission on the
	// SyncGroup.
	// TODO(m3b): The label should be read and write; can that be expressed?
	Join(ctx _gen_context.T, joiner NameIdentity, metaData JoinerMetaData, opts ..._gen_ipc.CallOpt) (reply SyncGroupInfo, err error)
	// Leave removes the joiner with the given name/identity from the
	// SyncGroup's Joiners map.
	//
	// Requires: this SyncGroup must exist;
	// the caller must assert the identity name.Identity.
	// (Thus, a device that Joined may Leave even if it would no longer
	// have permission to Join() the SyncGroup.)
	Leave(ctx _gen_context.T, name NameIdentity, opts ..._gen_ipc.CallOpt) (err error)
	// Eject is like Leave, but the caller must wield Admin
	// privilege on the group, and need not wield name.Identity.
	//
	// Requires: the SyncGroup must exist;
	// the caller must have admin permission on the SyncGroup.
	Eject(ctx _gen_context.T, name NameIdentity, opts ..._gen_ipc.CallOpt) (err error)
	// Destroy ejects all devices from the SyncGroup and removes it.
	// Devices that had joined will learn of this when their
	// SyncGroup object disappears.
	//
	// Requires: this SyncGroup must exist;
	// the caller must have admin permission on the SyncGroup.
	Destroy(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error)
	// Get returns the SyncGroupInfo for this SyncGroup.  For SyncGroups
	// with a large number of joiners, Get may return a subset of Joiners.
	//
	// Requires: this SyncGroup must exist;
	// the caller must have read permission on the SyncGroup.
	// TODO(m3b): This call may be removed when Watch is implemented.
	Get(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply SyncGroupInfo, err error)
	// SetConfig sets the Config field of this SyncGroup.
	//
	// Requires:  this SyncGroup must exist;
	// if non-empty, the eTag must match the value in the object;
	// the caller must have admin permission on the SyncGroup.
	SetConfig(ctx _gen_context.T, config SyncGroupConfig, eTag string, opts ..._gen_ipc.CallOpt) (err error)
}
type SyncGroupServer interface {
	_gen_ipc.UniversalServiceMethods
	SyncGroupServer_ExcludingUniversal
}

// SyncGroupServerService is the interface the server implements.
type SyncGroupServerService interface {

	// Object provides access control for Veyron objects.
	access.ObjectService
	// Create creates this SyncGroup with the given arguments, and if
	// joiner.Name!="", with {joiner, metaData} in its Joiners map.  It is
	// expected that acl will give Read and Write access to any device that
	// the administrator expects to join and sync; if the acl is empty, a
	// default ACL giving access only to the caller is used.  On success,
	// Create returns the SyncGroupInfo of the newly created object.
	//
	// Requires: this SyncGroup must not exist;
	// the caller must have write permission at the SyncGroup server;
	// the caller's identity must be a prefix of the SyncGroup's name.
	// Beware that for Create(), the access label is matched against a
	// server-wide ACL; for all other calls the access label is matched
	// against the object's ACL.
	Create(context _gen_ipc.ServerContext, createArgs SyncGroupConfig, rootOID storage.ID, joiner NameIdentity, metaData JoinerMetaData) (reply SyncGroupInfo, err error)
	// Join adds {joiner, metaData} to the SyncGroup's Joiners map and
	// returns the SyncGroupInfo for this SyncGroup.  The act of joining
	// allows other devices to find the caller, which is still required to
	// have read+write access on the SyncGroup to participate in
	// synchronization.  A device may call Join again with the same
	// NameIdentity in order to change metaData.
	// For SyncGroups with large numbers of joiners, Join may return
	// a subset of Joiners.
	//
	// Requires: this SyncGroup must exist;
	// the caller must have both read and write permission on the
	// SyncGroup.
	// TODO(m3b): The label should be read and write; can that be expressed?
	Join(context _gen_ipc.ServerContext, joiner NameIdentity, metaData JoinerMetaData) (reply SyncGroupInfo, err error)
	// Leave removes the joiner with the given name/identity from the
	// SyncGroup's Joiners map.
	//
	// Requires: this SyncGroup must exist;
	// the caller must assert the identity name.Identity.
	// (Thus, a device that Joined may Leave even if it would no longer
	// have permission to Join() the SyncGroup.)
	Leave(context _gen_ipc.ServerContext, name NameIdentity) (err error)
	// Eject is like Leave, but the caller must wield Admin
	// privilege on the group, and need not wield name.Identity.
	//
	// Requires: the SyncGroup must exist;
	// the caller must have admin permission on the SyncGroup.
	Eject(context _gen_ipc.ServerContext, name NameIdentity) (err error)
	// Destroy ejects all devices from the SyncGroup and removes it.
	// Devices that had joined will learn of this when their
	// SyncGroup object disappears.
	//
	// Requires: this SyncGroup must exist;
	// the caller must have admin permission on the SyncGroup.
	Destroy(context _gen_ipc.ServerContext) (err error)
	// Get returns the SyncGroupInfo for this SyncGroup.  For SyncGroups
	// with a large number of joiners, Get may return a subset of Joiners.
	//
	// Requires: this SyncGroup must exist;
	// the caller must have read permission on the SyncGroup.
	// TODO(m3b): This call may be removed when Watch is implemented.
	Get(context _gen_ipc.ServerContext) (reply SyncGroupInfo, err error)
	// SetConfig sets the Config field of this SyncGroup.
	//
	// Requires:  this SyncGroup must exist;
	// if non-empty, the eTag must match the value in the object;
	// the caller must have admin permission on the SyncGroup.
	SetConfig(context _gen_ipc.ServerContext, config SyncGroupConfig, eTag string) (err error)
}

// BindSyncGroupServer returns the client stub implementing the SyncGroupServer
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindSyncGroupServer(name string, opts ..._gen_ipc.BindOpt) (SyncGroupServer, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		// Do nothing.
	case 1:
		if clientOpt, ok := opts[0].(_gen_ipc.Client); opts[0] == nil || ok {
			client = clientOpt
		} else {
			return nil, _gen_vdlutil.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdlutil.ErrTooManyOptionsToBind
	}
	stub := &clientStubSyncGroupServer{defaultClient: client, name: name}
	stub.Object_ExcludingUniversal, _ = access.BindObject(name, client)

	return stub, nil
}

// NewServerSyncGroupServer creates a new server stub.
//
// It takes a regular server implementing the SyncGroupServerService
// interface, and returns a new server stub.
func NewServerSyncGroupServer(server SyncGroupServerService) interface{} {
	return &ServerStubSyncGroupServer{
		ServerStubObject: *access.NewServerObject(server).(*access.ServerStubObject),
		service:          server,
	}
}

// clientStubSyncGroupServer implements SyncGroupServer.
type clientStubSyncGroupServer struct {
	access.Object_ExcludingUniversal

	defaultClient _gen_ipc.Client
	name          string
}

func (__gen_c *clientStubSyncGroupServer) client(ctx _gen_context.T) _gen_ipc.Client {
	if __gen_c.defaultClient != nil {
		return __gen_c.defaultClient
	}
	return _gen_veyron2.RuntimeFromContext(ctx).Client()
}

func (__gen_c *clientStubSyncGroupServer) Create(ctx _gen_context.T, createArgs SyncGroupConfig, rootOID storage.ID, joiner NameIdentity, metaData JoinerMetaData, opts ..._gen_ipc.CallOpt) (reply SyncGroupInfo, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Create", []interface{}{createArgs, rootOID, joiner, metaData}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) Join(ctx _gen_context.T, joiner NameIdentity, metaData JoinerMetaData, opts ..._gen_ipc.CallOpt) (reply SyncGroupInfo, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Join", []interface{}{joiner, metaData}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) Leave(ctx _gen_context.T, name NameIdentity, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Leave", []interface{}{name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) Eject(ctx _gen_context.T, name NameIdentity, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Eject", []interface{}{name}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) Destroy(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Destroy", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) Get(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply SyncGroupInfo, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Get", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) SetConfig(ctx _gen_context.T, config SyncGroupConfig, eTag string, opts ..._gen_ipc.CallOpt) (err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "SetConfig", []interface{}{config, eTag}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) UnresolveStep(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply []string, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) Signature(ctx _gen_context.T, opts ..._gen_ipc.CallOpt) (reply _gen_ipc.ServiceSignature, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "Signature", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (__gen_c *clientStubSyncGroupServer) GetMethodTags(ctx _gen_context.T, method string, opts ..._gen_ipc.CallOpt) (reply []interface{}, err error) {
	var call _gen_ipc.Call
	if call, err = __gen_c.client(ctx).StartCall(ctx, __gen_c.name, "GetMethodTags", []interface{}{method}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubSyncGroupServer wraps a server that implements
// SyncGroupServerService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubSyncGroupServer struct {
	access.ServerStubObject

	service SyncGroupServerService
}

func (__gen_s *ServerStubSyncGroupServer) GetMethodTags(call _gen_ipc.ServerCall, method string) ([]interface{}, error) {
	// TODO(bprosnitz) GetMethodTags() will be replaces with Signature().
	// Note: This exhibits some weird behavior like returning a nil error if the method isn't found.
	// This will change when it is replaced with Signature().
	if resp, err := __gen_s.ServerStubObject.GetMethodTags(call, method); resp != nil || err != nil {
		return resp, err
	}
	switch method {
	case "Create":
		return []interface{}{security.Label(4)}, nil
	case "Join":
		return []interface{}{security.Label(4)}, nil
	case "Leave":
		return []interface{}{}, nil
	case "Eject":
		return []interface{}{security.Label(8)}, nil
	case "Destroy":
		return []interface{}{security.Label(8)}, nil
	case "Get":
		return []interface{}{security.Label(2)}, nil
	case "SetConfig":
		return []interface{}{security.Label(8)}, nil
	default:
		return nil, nil
	}
}

func (__gen_s *ServerStubSyncGroupServer) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Create"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "createArgs", Type: 72},
			{Name: "rootOID", Type: 74},
			{Name: "joiner", Type: 75},
			{Name: "metaData", Type: 76},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "sgInfo", Type: 79},
			{Name: "err", Type: 80},
		},
	}
	result.Methods["Destroy"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 80},
		},
	}
	result.Methods["Eject"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "name", Type: 75},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 80},
		},
	}
	result.Methods["Get"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "sgInfo", Type: 79},
			{Name: "err", Type: 80},
		},
	}
	result.Methods["Join"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "joiner", Type: 75},
			{Name: "metaData", Type: 76},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "sgInfo", Type: 79},
			{Name: "err", Type: 80},
		},
	}
	result.Methods["Leave"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "name", Type: 75},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 80},
		},
	}
	result.Methods["SetConfig"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "config", Type: 72},
			{Name: "eTag", Type: 3},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 80},
		},
	}

	result.TypeDefs = []_gen_vdlutil.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "anydata", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x3, Elem: 0x41, Name: "", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x3, Name: "veyron2/security.PrincipalPattern", Tags: []string(nil)}, _gen_wiretype.NamedPrimitiveType{Type: 0x34, Name: "veyron2/security.LabelSet", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x43, Elem: 0x44, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x45, Name: "Principals"},
			},
			"veyron2/security.Entries", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x46, Name: "In"},
				_gen_wiretype.FieldType{Type: 0x46, Name: "NotIn"},
			},
			"veyron2/security.ACL", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Desc"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "PathPatterns"},
				_gen_wiretype.FieldType{Type: 0x42, Name: "Options"},
				_gen_wiretype.FieldType{Type: 0x47, Name: "ACL"},
				_gen_wiretype.FieldType{Type: 0x3d, Name: "MountTables"},
			},
			"veyron/services/syncgroup.SyncGroupConfig", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x32, Name: "byte", Tags: []string(nil)}, _gen_wiretype.ArrayType{Elem: 0x49, Len: 0x10, Name: "veyron2/storage.ID", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "Identity"},
			},
			"veyron/services/syncgroup.NameIdentity", []string(nil)},
		_gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x24, Name: "SyncPriority"},
			},
			"veyron/services/syncgroup.JoinerMetaData", []string(nil)},
		_gen_wiretype.ArrayType{Elem: 0x49, Len: 0x10, Name: "veyron/services/syncgroup.ID", Tags: []string(nil)}, _gen_wiretype.MapType{Key: 0x4b, Elem: 0x4c, Name: "", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x3, Name: "Name"},
				_gen_wiretype.FieldType{Type: 0x48, Name: "Config"},
				_gen_wiretype.FieldType{Type: 0x4a, Name: "RootOID"},
				_gen_wiretype.FieldType{Type: 0x3, Name: "ETag"},
				_gen_wiretype.FieldType{Type: 0x4d, Name: "SGOID"},
				_gen_wiretype.FieldType{Type: 0x4e, Name: "Joiners"},
			},
			"veyron/services/syncgroup.SyncGroupInfo", []string(nil)},
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}
	var ss _gen_ipc.ServiceSignature
	var firstAdded int
	ss, _ = __gen_s.ServerStubObject.Signature(call)
	firstAdded = len(result.TypeDefs)
	for k, v := range ss.Methods {
		for i, _ := range v.InArgs {
			if v.InArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.InArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		for i, _ := range v.OutArgs {
			if v.OutArgs[i].Type >= _gen_wiretype.TypeIDFirst {
				v.OutArgs[i].Type += _gen_wiretype.TypeID(firstAdded)
			}
		}
		if v.InStream >= _gen_wiretype.TypeIDFirst {
			v.InStream += _gen_wiretype.TypeID(firstAdded)
		}
		if v.OutStream >= _gen_wiretype.TypeIDFirst {
			v.OutStream += _gen_wiretype.TypeID(firstAdded)
		}
		result.Methods[k] = v
	}
	//TODO(bprosnitz) combine type definitions from embeded interfaces in a way that doesn't cause duplication.
	for _, d := range ss.TypeDefs {
		switch wt := d.(type) {
		case _gen_wiretype.SliceType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.ArrayType:
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.MapType:
			if wt.Key >= _gen_wiretype.TypeIDFirst {
				wt.Key += _gen_wiretype.TypeID(firstAdded)
			}
			if wt.Elem >= _gen_wiretype.TypeIDFirst {
				wt.Elem += _gen_wiretype.TypeID(firstAdded)
			}
			d = wt
		case _gen_wiretype.StructType:
			for i, fld := range wt.Fields {
				if fld.Type >= _gen_wiretype.TypeIDFirst {
					wt.Fields[i].Type += _gen_wiretype.TypeID(firstAdded)
				}
			}
			d = wt
			// NOTE: other types are missing, but we are upgrading anyways.
		}
		result.TypeDefs = append(result.TypeDefs, d)
	}

	return result, nil
}

func (__gen_s *ServerStubSyncGroupServer) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := __gen_s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubSyncGroupServer) Create(call _gen_ipc.ServerCall, createArgs SyncGroupConfig, rootOID storage.ID, joiner NameIdentity, metaData JoinerMetaData) (reply SyncGroupInfo, err error) {
	reply, err = __gen_s.service.Create(call, createArgs, rootOID, joiner, metaData)
	return
}

func (__gen_s *ServerStubSyncGroupServer) Join(call _gen_ipc.ServerCall, joiner NameIdentity, metaData JoinerMetaData) (reply SyncGroupInfo, err error) {
	reply, err = __gen_s.service.Join(call, joiner, metaData)
	return
}

func (__gen_s *ServerStubSyncGroupServer) Leave(call _gen_ipc.ServerCall, name NameIdentity) (err error) {
	err = __gen_s.service.Leave(call, name)
	return
}

func (__gen_s *ServerStubSyncGroupServer) Eject(call _gen_ipc.ServerCall, name NameIdentity) (err error) {
	err = __gen_s.service.Eject(call, name)
	return
}

func (__gen_s *ServerStubSyncGroupServer) Destroy(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.Destroy(call)
	return
}

func (__gen_s *ServerStubSyncGroupServer) Get(call _gen_ipc.ServerCall) (reply SyncGroupInfo, err error) {
	reply, err = __gen_s.service.Get(call)
	return
}

func (__gen_s *ServerStubSyncGroupServer) SetConfig(call _gen_ipc.ServerCall, config SyncGroupConfig, eTag string) (err error) {
	err = __gen_s.service.SetConfig(call, config, eTag)
	return
}
