// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: wire.vdl

// Package agent defines an interface to keep a private key in memory, and for
// clients to have access to the private key.
//
// Protocol
//
// The agent starts processes with the VEYRON_AGENT_FD set to one end of a
// unix domain socket. To connect to the agent, a client should create
// a unix domain socket pair. Then send one end of the socket to the agent
// with 1 byte of data. The agent will then serve the Agent service on
// the received socket, using SecurityNone.
//
// The agent also supports an optional mode where it can manage multiple principals.
// Typically this is only used by Device Manager. In this mode, VEYRON_AGENT_FD
// will be 3, and there will be another socket at fd 4.
// Creating a new principal is similar to connecting to to agent: create a socket
// pair and send one end on fd 4 with 1 byte of data.
// Set the data to 1 to request the principal only be stored in memory.
// The agent will create a new principal and respond with a principal handle on fd 4.
// To connect using a previously created principal, create a socket pair and send
// one end with the principal handle as data on fd 4. The agent will not send a
// response on fd 4.
// In either, you can use the normal process to connect to an agent over the
// other end of the pair. Typically you would pass the other end to a child
// process and set VEYRON_AGENT_FD so it knows to connect.
//
// The protocol also has limited support for caching: A client can
// request notification when any other client modifies the principal so it
// can flush the cache. See NotifyWhenChanged for details.
package agent

import (
	"fmt"
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security"
	"v.io/v23/vdl"
	"v.io/v23/verror"
)

type ConnInfo struct {
	MinVersion int32
	MaxVersion int32
}

func (ConnInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/agent.ConnInfo"`
}) {
}

func (m *ConnInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_wire_v_io_x_ref_services_agent_ConnInfo == nil || __VDLTypewire0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("MinVersion")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromInt(int64(m.MinVersion), vdl.Int32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("MaxVersion")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromInt(int64(m.MaxVersion), vdl.Int32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *ConnInfo) MakeVDLTarget() vdl.Target {
	return &ConnInfoTarget{Value: m}
}

type ConnInfoTarget struct {
	Value *ConnInfo
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *ConnInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_wire_v_io_x_ref_services_agent_ConnInfo) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_wire_v_io_x_ref_services_agent_ConnInfo)
	}
	return t, nil
}
func (t *ConnInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "MinVersion":
		val, err := &vdl.Int32Target{Value: &t.Value.MinVersion}, error(nil)
		return nil, val, err
	case "MaxVersion":
		val, err := &vdl.Int32Target{Value: &t.Value.MaxVersion}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_wire_v_io_x_ref_services_agent_ConnInfo)
	}
}
func (t *ConnInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *ConnInfoTarget) FinishFields(_ vdl.FieldsTarget) error {
	return nil
}

type RpcRequest struct {
	Id      uint64
	Method  string
	NumArgs uint32
}

func (RpcRequest) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/agent.RpcRequest"`
}) {
}

func (m *RpcRequest) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_wire_v_io_x_ref_services_agent_RpcRequest == nil || __VDLTypewire1 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.Id), vdl.Uint64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Method")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget5.FromString(string(m.Method), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("NumArgs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget7.FromUint(uint64(m.NumArgs), vdl.Uint32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *RpcRequest) MakeVDLTarget() vdl.Target {
	return &RpcRequestTarget{Value: m}
}

type RpcRequestTarget struct {
	Value *RpcRequest
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *RpcRequestTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_wire_v_io_x_ref_services_agent_RpcRequest) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_wire_v_io_x_ref_services_agent_RpcRequest)
	}
	return t, nil
}
func (t *RpcRequestTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		val, err := &vdl.Uint64Target{Value: &t.Value.Id}, error(nil)
		return nil, val, err
	case "Method":
		val, err := &vdl.StringTarget{Value: &t.Value.Method}, error(nil)
		return nil, val, err
	case "NumArgs":
		val, err := &vdl.Uint32Target{Value: &t.Value.NumArgs}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_wire_v_io_x_ref_services_agent_RpcRequest)
	}
}
func (t *RpcRequestTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *RpcRequestTarget) FinishFields(_ vdl.FieldsTarget) error {
	return nil
}

type RpcResponse struct {
	Id      uint64
	Err     error
	NumArgs uint32
}

func (RpcResponse) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/agent.RpcResponse"`
}) {
}

func (m *RpcResponse) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_wire_v_io_x_ref_services_agent_RpcResponse == nil || __VDLTypewire2 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Id")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromUint(uint64(m.Id), vdl.Uint64Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Err")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Err == nil {
			if err := fieldTarget5.FromNil(vdl.ErrorType); err != nil {
				return err
			}
		} else {
			var wireError6 vdl.WireError
			if err := verror.WireFromNative(&wireError6, m.Err); err != nil {
				return err
			}
			if err := wireError6.FillVDLTarget(fieldTarget5, vdl.ErrorType); err != nil {
				return err
			}

		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("NumArgs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget8.FromUint(uint64(m.NumArgs), vdl.Uint32Type); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *RpcResponse) MakeVDLTarget() vdl.Target {
	return &RpcResponseTarget{Value: m}
}

type RpcResponseTarget struct {
	Value *RpcResponse
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *RpcResponseTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_wire_v_io_x_ref_services_agent_RpcResponse) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_wire_v_io_x_ref_services_agent_RpcResponse)
	}
	return t, nil
}
func (t *RpcResponseTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Id":
		val, err := &vdl.Uint64Target{Value: &t.Value.Id}, error(nil)
		return nil, val, err
	case "Err":
		val, err := &verror.ErrorTarget{Value: &t.Value.Err}, error(nil)
		return nil, val, err
	case "NumArgs":
		val, err := &vdl.Uint32Target{Value: &t.Value.NumArgs}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_wire_v_io_x_ref_services_agent_RpcResponse)
	}
}
func (t *RpcResponseTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *RpcResponseTarget) FinishFields(_ vdl.FieldsTarget) error {
	return nil
}

type (
	// RpcMessage represents any single field of the RpcMessage union type.
	RpcMessage interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the RpcMessage union type.
		__VDLReflect(__RpcMessageReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// RpcMessageReq represents field Req of the RpcMessage union type.
	RpcMessageReq struct{ Value RpcRequest }
	// RpcMessageResp represents field Resp of the RpcMessage union type.
	RpcMessageResp struct{ Value RpcResponse }
	// __RpcMessageReflect describes the RpcMessage union type.
	__RpcMessageReflect struct {
		Name  string `vdl:"v.io/x/ref/services/agent.RpcMessage"`
		Type  RpcMessage
		Union struct {
			Req  RpcMessageReq
			Resp RpcMessageResp
		}
	}
)

func (x RpcMessageReq) Index() int                       { return 0 }
func (x RpcMessageReq) Interface() interface{}           { return x.Value }
func (x RpcMessageReq) Name() string                     { return "Req" }
func (x RpcMessageReq) __VDLReflect(__RpcMessageReflect) {}

func (m RpcMessageReq) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(__VDLType_wire_v_io_x_ref_services_agent_RpcMessage)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Req")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_wire_v_io_x_ref_services_agent_RpcRequest); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m RpcMessageReq) MakeVDLTarget() vdl.Target {
	return nil
}

func (x RpcMessageResp) Index() int                       { return 1 }
func (x RpcMessageResp) Interface() interface{}           { return x.Value }
func (x RpcMessageResp) Name() string                     { return "Resp" }
func (x RpcMessageResp) __VDLReflect(__RpcMessageReflect) {}

func (m RpcMessageResp) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(__VDLType_wire_v_io_x_ref_services_agent_RpcMessage)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Resp")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_wire_v_io_x_ref_services_agent_RpcResponse); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m RpcMessageResp) MakeVDLTarget() vdl.Target {
	return nil
}

func init() {
	vdl.Register((*ConnInfo)(nil))
	vdl.Register((*RpcRequest)(nil))
	vdl.Register((*RpcResponse)(nil))
	vdl.Register((*RpcMessage)(nil))
}

var __VDLTypewire0 *vdl.Type = vdl.TypeOf((*ConnInfo)(nil))
var __VDLTypewire1 *vdl.Type = vdl.TypeOf((*RpcRequest)(nil))
var __VDLTypewire2 *vdl.Type = vdl.TypeOf((*RpcResponse)(nil))
var __VDLType_wire_v_io_x_ref_services_agent_ConnInfo *vdl.Type = vdl.TypeOf(ConnInfo{})
var __VDLType_wire_v_io_x_ref_services_agent_RpcMessage *vdl.Type = vdl.TypeOf(RpcMessage(RpcMessageReq{RpcRequest{}}))
var __VDLType_wire_v_io_x_ref_services_agent_RpcRequest *vdl.Type = vdl.TypeOf(RpcRequest{})
var __VDLType_wire_v_io_x_ref_services_agent_RpcResponse *vdl.Type = vdl.TypeOf(RpcResponse{})

func __VDLEnsureNativeBuilt_wire() {
}

// AgentClientMethods is the client interface
// containing Agent methods.
type AgentClientMethods interface {
	Bless(_ *context.T, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat, _ ...rpc.CallOpt) (security.Blessings, error)
	BlessSelf(_ *context.T, name string, caveats []security.Caveat, _ ...rpc.CallOpt) (security.Blessings, error)
	Sign(_ *context.T, message []byte, _ ...rpc.CallOpt) (security.Signature, error)
	MintDischarge(_ *context.T, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat, _ ...rpc.CallOpt) (security.Discharge, error)
	PublicKey(*context.T, ...rpc.CallOpt) ([]byte, error)
	BlessingStoreSet(_ *context.T, blessings security.Blessings, forPeers security.BlessingPattern, _ ...rpc.CallOpt) (security.Blessings, error)
	BlessingStoreForPeer(_ *context.T, peerBlessings []string, _ ...rpc.CallOpt) (security.Blessings, error)
	BlessingStoreSetDefault(_ *context.T, blessings security.Blessings, _ ...rpc.CallOpt) error
	BlessingStoreDefault(*context.T, ...rpc.CallOpt) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, ...rpc.CallOpt) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, ...rpc.CallOpt) (string, error)
	BlessingStoreCacheDischarge(_ *context.T, discharge security.Discharge, caveat security.Caveat, impetus security.DischargeImpetus, _ ...rpc.CallOpt) error
	BlessingStoreClearDischarges(_ *context.T, discharges []security.Discharge, _ ...rpc.CallOpt) error
	BlessingStoreDischarge(_ *context.T, caveat security.Caveat, impetus security.DischargeImpetus, _ ...rpc.CallOpt) (wd security.Discharge, _ error)
	BlessingRootsAdd(_ *context.T, root []byte, pattern security.BlessingPattern, _ ...rpc.CallOpt) error
	BlessingRootsRecognized(_ *context.T, root []byte, blessing string, _ ...rpc.CallOpt) error
	BlessingRootsDump(*context.T, ...rpc.CallOpt) (map[security.BlessingPattern][][]byte, error)
	BlessingRootsDebugString(*context.T, ...rpc.CallOpt) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, ...rpc.CallOpt) (AgentNotifyWhenChangedClientCall, error)
}

// AgentClientStub adds universal methods to AgentClientMethods.
type AgentClientStub interface {
	AgentClientMethods
	rpc.UniversalServiceMethods
}

// AgentClient returns a client stub for Agent.
func AgentClient(name string) AgentClientStub {
	return implAgentClientStub{name}
}

type implAgentClientStub struct {
	name string
}

func (c implAgentClientStub) Bless(ctx *context.T, i0 []byte, i1 security.Blessings, i2 string, i3 security.Caveat, i4 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Bless", []interface{}{i0, i1, i2, i3, i4}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessSelf(ctx *context.T, i0 string, i1 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessSelf", []interface{}{i0, i1}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) Sign(ctx *context.T, i0 []byte, opts ...rpc.CallOpt) (o0 security.Signature, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Sign", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) MintDischarge(ctx *context.T, i0 security.Caveat, i1 security.Caveat, i2 []security.Caveat, opts ...rpc.CallOpt) (o0 security.Discharge, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "MintDischarge", []interface{}{i0, i1, i2}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) PublicKey(ctx *context.T, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "PublicKey", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreSet(ctx *context.T, i0 security.Blessings, i1 security.BlessingPattern, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreSet", []interface{}{i0, i1}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreForPeer(ctx *context.T, i0 []string, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreForPeer", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreSetDefault(ctx *context.T, i0 security.Blessings, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreSetDefault", []interface{}{i0}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreDefault(ctx *context.T, opts ...rpc.CallOpt) (o0 security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreDefault", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStorePeerBlessings(ctx *context.T, opts ...rpc.CallOpt) (o0 map[security.BlessingPattern]security.Blessings, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStorePeerBlessings", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreDebugString(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreDebugString", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreCacheDischarge(ctx *context.T, i0 security.Discharge, i1 security.Caveat, i2 security.DischargeImpetus, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreCacheDischarge", []interface{}{i0, i1, i2}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreClearDischarges(ctx *context.T, i0 []security.Discharge, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreClearDischarges", []interface{}{i0}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingStoreDischarge(ctx *context.T, i0 security.Caveat, i1 security.DischargeImpetus, opts ...rpc.CallOpt) (o0 security.Discharge, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingStoreDischarge", []interface{}{i0, i1}, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsAdd(ctx *context.T, i0 []byte, i1 security.BlessingPattern, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsAdd", []interface{}{i0, i1}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsRecognized(ctx *context.T, i0 []byte, i1 string, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsRecognized", []interface{}{i0, i1}, nil, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsDump(ctx *context.T, opts ...rpc.CallOpt) (o0 map[security.BlessingPattern][][]byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsDump", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) BlessingRootsDebugString(ctx *context.T, opts ...rpc.CallOpt) (o0 string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "BlessingRootsDebugString", nil, []interface{}{&o0}, opts...)
	return
}

func (c implAgentClientStub) NotifyWhenChanged(ctx *context.T, opts ...rpc.CallOpt) (ocall AgentNotifyWhenChangedClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "NotifyWhenChanged", nil, opts...); err != nil {
		return
	}
	ocall = &implAgentNotifyWhenChangedClientCall{ClientCall: call}
	return
}

// AgentNotifyWhenChangedClientStream is the client stream for Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedClientStream interface {
	// RecvStream returns the receiver side of the Agent.NotifyWhenChanged client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() bool
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// AgentNotifyWhenChangedClientCall represents the call returned from Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedClientCall interface {
	AgentNotifyWhenChangedClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implAgentNotifyWhenChangedClientCall struct {
	rpc.ClientCall
	valRecv bool
	errRecv error
}

func (c *implAgentNotifyWhenChangedClientCall) RecvStream() interface {
	Advance() bool
	Value() bool
	Err() error
} {
	return implAgentNotifyWhenChangedClientCallRecv{c}
}

type implAgentNotifyWhenChangedClientCallRecv struct {
	c *implAgentNotifyWhenChangedClientCall
}

func (c implAgentNotifyWhenChangedClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implAgentNotifyWhenChangedClientCallRecv) Value() bool {
	return c.c.valRecv
}
func (c implAgentNotifyWhenChangedClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implAgentNotifyWhenChangedClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// AgentServerMethods is the interface a server writer
// implements for Agent.
type AgentServerMethods interface {
	Bless(_ *context.T, _ rpc.ServerCall, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat) (security.Blessings, error)
	BlessSelf(_ *context.T, _ rpc.ServerCall, name string, caveats []security.Caveat) (security.Blessings, error)
	Sign(_ *context.T, _ rpc.ServerCall, message []byte) (security.Signature, error)
	MintDischarge(_ *context.T, _ rpc.ServerCall, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat) (security.Discharge, error)
	PublicKey(*context.T, rpc.ServerCall) ([]byte, error)
	BlessingStoreSet(_ *context.T, _ rpc.ServerCall, blessings security.Blessings, forPeers security.BlessingPattern) (security.Blessings, error)
	BlessingStoreForPeer(_ *context.T, _ rpc.ServerCall, peerBlessings []string) (security.Blessings, error)
	BlessingStoreSetDefault(_ *context.T, _ rpc.ServerCall, blessings security.Blessings) error
	BlessingStoreDefault(*context.T, rpc.ServerCall) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, rpc.ServerCall) (string, error)
	BlessingStoreCacheDischarge(_ *context.T, _ rpc.ServerCall, discharge security.Discharge, caveat security.Caveat, impetus security.DischargeImpetus) error
	BlessingStoreClearDischarges(_ *context.T, _ rpc.ServerCall, discharges []security.Discharge) error
	BlessingStoreDischarge(_ *context.T, _ rpc.ServerCall, caveat security.Caveat, impetus security.DischargeImpetus) (wd security.Discharge, _ error)
	BlessingRootsAdd(_ *context.T, _ rpc.ServerCall, root []byte, pattern security.BlessingPattern) error
	BlessingRootsRecognized(_ *context.T, _ rpc.ServerCall, root []byte, blessing string) error
	BlessingRootsDump(*context.T, rpc.ServerCall) (map[security.BlessingPattern][][]byte, error)
	BlessingRootsDebugString(*context.T, rpc.ServerCall) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, AgentNotifyWhenChangedServerCall) error
}

// AgentServerStubMethods is the server interface containing
// Agent methods, as expected by rpc.Server.
// The only difference between this interface and AgentServerMethods
// is the streaming methods.
type AgentServerStubMethods interface {
	Bless(_ *context.T, _ rpc.ServerCall, key []byte, wit security.Blessings, extension string, caveat security.Caveat, additionalCaveats []security.Caveat) (security.Blessings, error)
	BlessSelf(_ *context.T, _ rpc.ServerCall, name string, caveats []security.Caveat) (security.Blessings, error)
	Sign(_ *context.T, _ rpc.ServerCall, message []byte) (security.Signature, error)
	MintDischarge(_ *context.T, _ rpc.ServerCall, forCaveat security.Caveat, caveatOnDischarge security.Caveat, additionalCaveatsOnDischarge []security.Caveat) (security.Discharge, error)
	PublicKey(*context.T, rpc.ServerCall) ([]byte, error)
	BlessingStoreSet(_ *context.T, _ rpc.ServerCall, blessings security.Blessings, forPeers security.BlessingPattern) (security.Blessings, error)
	BlessingStoreForPeer(_ *context.T, _ rpc.ServerCall, peerBlessings []string) (security.Blessings, error)
	BlessingStoreSetDefault(_ *context.T, _ rpc.ServerCall, blessings security.Blessings) error
	BlessingStoreDefault(*context.T, rpc.ServerCall) (security.Blessings, error)
	BlessingStorePeerBlessings(*context.T, rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error)
	BlessingStoreDebugString(*context.T, rpc.ServerCall) (string, error)
	BlessingStoreCacheDischarge(_ *context.T, _ rpc.ServerCall, discharge security.Discharge, caveat security.Caveat, impetus security.DischargeImpetus) error
	BlessingStoreClearDischarges(_ *context.T, _ rpc.ServerCall, discharges []security.Discharge) error
	BlessingStoreDischarge(_ *context.T, _ rpc.ServerCall, caveat security.Caveat, impetus security.DischargeImpetus) (wd security.Discharge, _ error)
	BlessingRootsAdd(_ *context.T, _ rpc.ServerCall, root []byte, pattern security.BlessingPattern) error
	BlessingRootsRecognized(_ *context.T, _ rpc.ServerCall, root []byte, blessing string) error
	BlessingRootsDump(*context.T, rpc.ServerCall) (map[security.BlessingPattern][][]byte, error)
	BlessingRootsDebugString(*context.T, rpc.ServerCall) (string, error)
	// Clients using caching should call NotifyWhenChanged upon connecting to
	// the server. The server will stream back values whenever the client should
	// flush the cache. The streamed value is arbitrary, simply flush whenever
	// recieving a new item.
	NotifyWhenChanged(*context.T, *AgentNotifyWhenChangedServerCallStub) error
}

// AgentServerStub adds universal methods to AgentServerStubMethods.
type AgentServerStub interface {
	AgentServerStubMethods
	// Describe the Agent interfaces.
	Describe__() []rpc.InterfaceDesc
}

// AgentServer returns a server stub for Agent.
// It converts an implementation of AgentServerMethods into
// an object that may be used by rpc.Server.
func AgentServer(impl AgentServerMethods) AgentServerStub {
	stub := implAgentServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implAgentServerStub struct {
	impl AgentServerMethods
	gs   *rpc.GlobState
}

func (s implAgentServerStub) Bless(ctx *context.T, call rpc.ServerCall, i0 []byte, i1 security.Blessings, i2 string, i3 security.Caveat, i4 []security.Caveat) (security.Blessings, error) {
	return s.impl.Bless(ctx, call, i0, i1, i2, i3, i4)
}

func (s implAgentServerStub) BlessSelf(ctx *context.T, call rpc.ServerCall, i0 string, i1 []security.Caveat) (security.Blessings, error) {
	return s.impl.BlessSelf(ctx, call, i0, i1)
}

func (s implAgentServerStub) Sign(ctx *context.T, call rpc.ServerCall, i0 []byte) (security.Signature, error) {
	return s.impl.Sign(ctx, call, i0)
}

func (s implAgentServerStub) MintDischarge(ctx *context.T, call rpc.ServerCall, i0 security.Caveat, i1 security.Caveat, i2 []security.Caveat) (security.Discharge, error) {
	return s.impl.MintDischarge(ctx, call, i0, i1, i2)
}

func (s implAgentServerStub) PublicKey(ctx *context.T, call rpc.ServerCall) ([]byte, error) {
	return s.impl.PublicKey(ctx, call)
}

func (s implAgentServerStub) BlessingStoreSet(ctx *context.T, call rpc.ServerCall, i0 security.Blessings, i1 security.BlessingPattern) (security.Blessings, error) {
	return s.impl.BlessingStoreSet(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingStoreForPeer(ctx *context.T, call rpc.ServerCall, i0 []string) (security.Blessings, error) {
	return s.impl.BlessingStoreForPeer(ctx, call, i0)
}

func (s implAgentServerStub) BlessingStoreSetDefault(ctx *context.T, call rpc.ServerCall, i0 security.Blessings) error {
	return s.impl.BlessingStoreSetDefault(ctx, call, i0)
}

func (s implAgentServerStub) BlessingStoreDefault(ctx *context.T, call rpc.ServerCall) (security.Blessings, error) {
	return s.impl.BlessingStoreDefault(ctx, call)
}

func (s implAgentServerStub) BlessingStorePeerBlessings(ctx *context.T, call rpc.ServerCall) (map[security.BlessingPattern]security.Blessings, error) {
	return s.impl.BlessingStorePeerBlessings(ctx, call)
}

func (s implAgentServerStub) BlessingStoreDebugString(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.BlessingStoreDebugString(ctx, call)
}

func (s implAgentServerStub) BlessingStoreCacheDischarge(ctx *context.T, call rpc.ServerCall, i0 security.Discharge, i1 security.Caveat, i2 security.DischargeImpetus) error {
	return s.impl.BlessingStoreCacheDischarge(ctx, call, i0, i1, i2)
}

func (s implAgentServerStub) BlessingStoreClearDischarges(ctx *context.T, call rpc.ServerCall, i0 []security.Discharge) error {
	return s.impl.BlessingStoreClearDischarges(ctx, call, i0)
}

func (s implAgentServerStub) BlessingStoreDischarge(ctx *context.T, call rpc.ServerCall, i0 security.Caveat, i1 security.DischargeImpetus) (security.Discharge, error) {
	return s.impl.BlessingStoreDischarge(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsAdd(ctx *context.T, call rpc.ServerCall, i0 []byte, i1 security.BlessingPattern) error {
	return s.impl.BlessingRootsAdd(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsRecognized(ctx *context.T, call rpc.ServerCall, i0 []byte, i1 string) error {
	return s.impl.BlessingRootsRecognized(ctx, call, i0, i1)
}

func (s implAgentServerStub) BlessingRootsDump(ctx *context.T, call rpc.ServerCall) (map[security.BlessingPattern][][]byte, error) {
	return s.impl.BlessingRootsDump(ctx, call)
}

func (s implAgentServerStub) BlessingRootsDebugString(ctx *context.T, call rpc.ServerCall) (string, error) {
	return s.impl.BlessingRootsDebugString(ctx, call)
}

func (s implAgentServerStub) NotifyWhenChanged(ctx *context.T, call *AgentNotifyWhenChangedServerCallStub) error {
	return s.impl.NotifyWhenChanged(ctx, call)
}

func (s implAgentServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implAgentServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{AgentDesc}
}

// AgentDesc describes the Agent interface.
var AgentDesc rpc.InterfaceDesc = descAgent

// descAgent hides the desc to keep godoc clean.
var descAgent = rpc.InterfaceDesc{
	Name:    "Agent",
	PkgPath: "v.io/x/ref/services/agent",
	Methods: []rpc.MethodDesc{
		{
			Name: "Bless",
			InArgs: []rpc.ArgDesc{
				{"key", ``},               // []byte
				{"wit", ``},               // security.Blessings
				{"extension", ``},         // string
				{"caveat", ``},            // security.Caveat
				{"additionalCaveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessSelf",
			InArgs: []rpc.ArgDesc{
				{"name", ``},    // string
				{"caveats", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "Sign",
			InArgs: []rpc.ArgDesc{
				{"message", ``}, // []byte
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Signature
			},
		},
		{
			Name: "MintDischarge",
			InArgs: []rpc.ArgDesc{
				{"forCaveat", ``},                    // security.Caveat
				{"caveatOnDischarge", ``},            // security.Caveat
				{"additionalCaveatsOnDischarge", ``}, // []security.Caveat
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Discharge
			},
		},
		{
			Name: "PublicKey",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []byte
			},
		},
		{
			Name: "BlessingStoreSet",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
				{"forPeers", ``},  // security.BlessingPattern
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreForPeer",
			InArgs: []rpc.ArgDesc{
				{"peerBlessings", ``}, // []string
			},
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreSetDefault",
			InArgs: []rpc.ArgDesc{
				{"blessings", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStoreDefault",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // security.Blessings
			},
		},
		{
			Name: "BlessingStorePeerBlessings",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[security.BlessingPattern]security.Blessings
			},
		},
		{
			Name: "BlessingStoreDebugString",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "BlessingStoreCacheDischarge",
			InArgs: []rpc.ArgDesc{
				{"discharge", ``}, // security.Discharge
				{"caveat", ``},    // security.Caveat
				{"impetus", ``},   // security.DischargeImpetus
			},
		},
		{
			Name: "BlessingStoreClearDischarges",
			InArgs: []rpc.ArgDesc{
				{"discharges", ``}, // []security.Discharge
			},
		},
		{
			Name: "BlessingStoreDischarge",
			InArgs: []rpc.ArgDesc{
				{"caveat", ``},  // security.Caveat
				{"impetus", ``}, // security.DischargeImpetus
			},
			OutArgs: []rpc.ArgDesc{
				{"wd", ``}, // security.Discharge
			},
		},
		{
			Name: "BlessingRootsAdd",
			InArgs: []rpc.ArgDesc{
				{"root", ``},    // []byte
				{"pattern", ``}, // security.BlessingPattern
			},
		},
		{
			Name: "BlessingRootsRecognized",
			InArgs: []rpc.ArgDesc{
				{"root", ``},     // []byte
				{"blessing", ``}, // string
			},
		},
		{
			Name: "BlessingRootsDump",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // map[security.BlessingPattern][][]byte
			},
		},
		{
			Name: "BlessingRootsDebugString",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // string
			},
		},
		{
			Name: "NotifyWhenChanged",
			Doc:  "// Clients using caching should call NotifyWhenChanged upon connecting to\n// the server. The server will stream back values whenever the client should\n// flush the cache. The streamed value is arbitrary, simply flush whenever\n// recieving a new item.",
		},
	},
}

// AgentNotifyWhenChangedServerStream is the server stream for Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedServerStream interface {
	// SendStream returns the send side of the Agent.NotifyWhenChanged server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item bool) error
	}
}

// AgentNotifyWhenChangedServerCall represents the context passed to Agent.NotifyWhenChanged.
type AgentNotifyWhenChangedServerCall interface {
	rpc.ServerCall
	AgentNotifyWhenChangedServerStream
}

// AgentNotifyWhenChangedServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements AgentNotifyWhenChangedServerCall.
type AgentNotifyWhenChangedServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes AgentNotifyWhenChangedServerCallStub from rpc.StreamServerCall.
func (s *AgentNotifyWhenChangedServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the Agent.NotifyWhenChanged server stream.
func (s *AgentNotifyWhenChangedServerCallStub) SendStream() interface {
	Send(item bool) error
} {
	return implAgentNotifyWhenChangedServerCallSend{s}
}

type implAgentNotifyWhenChangedServerCallSend struct {
	s *AgentNotifyWhenChangedServerCallStub
}

func (s implAgentNotifyWhenChangedServerCallSend) Send(item bool) error {
	return s.s.Send(item)
}
