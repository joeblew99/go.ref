// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: root.vdl

// Package iberoot defines an interface for requesting private keys for
// specific blessings in a blessings-based encryption scheme.
package iberoot

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/x/ref/lib/security/bcrypter"
)

func __VDLEnsureNativeBuilt_root() {
}

// RootClientMethods is the client interface
// containing Root methods.
//
// Root is an interface for requesting private keys for blessings.
//
// The keys are extracted in a blessings-based encryption scheme, which in
// turn is based on an identity-based encryption (IBE) scheme (e.g., the BB1
// IBE scheme).
type RootClientMethods interface {
	// SeekPrivateKeys creates and returns private keys for blessings
	// presented by the calling principal. The blessings must be from
	// an identity provider recognized by this service.
	//
	// The extracted private keys can be used to decrypt any ciphertext
	// encrypted for a pattern matched by the presented blessings.
	SeekPrivateKeys(*context.T, ...rpc.CallOpt) ([]bcrypter.WirePrivateKey, error)
	// Params returns the public encryption parameters of this service.
	Params(*context.T, ...rpc.CallOpt) (bcrypter.WireParams, error)
}

// RootClientStub adds universal methods to RootClientMethods.
type RootClientStub interface {
	RootClientMethods
	rpc.UniversalServiceMethods
}

// RootClient returns a client stub for Root.
func RootClient(name string) RootClientStub {
	return implRootClientStub{name}
}

type implRootClientStub struct {
	name string
}

func (c implRootClientStub) SeekPrivateKeys(ctx *context.T, opts ...rpc.CallOpt) (o0 []bcrypter.WirePrivateKey, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "SeekPrivateKeys", nil, []interface{}{&o0}, opts...)
	return
}

func (c implRootClientStub) Params(ctx *context.T, opts ...rpc.CallOpt) (o0 bcrypter.WireParams, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Params", nil, []interface{}{&o0}, opts...)
	return
}

// RootServerMethods is the interface a server writer
// implements for Root.
//
// Root is an interface for requesting private keys for blessings.
//
// The keys are extracted in a blessings-based encryption scheme, which in
// turn is based on an identity-based encryption (IBE) scheme (e.g., the BB1
// IBE scheme).
type RootServerMethods interface {
	// SeekPrivateKeys creates and returns private keys for blessings
	// presented by the calling principal. The blessings must be from
	// an identity provider recognized by this service.
	//
	// The extracted private keys can be used to decrypt any ciphertext
	// encrypted for a pattern matched by the presented blessings.
	SeekPrivateKeys(*context.T, rpc.ServerCall) ([]bcrypter.WirePrivateKey, error)
	// Params returns the public encryption parameters of this service.
	Params(*context.T, rpc.ServerCall) (bcrypter.WireParams, error)
}

// RootServerStubMethods is the server interface containing
// Root methods, as expected by rpc.Server.
// There is no difference between this interface and RootServerMethods
// since there are no streaming methods.
type RootServerStubMethods RootServerMethods

// RootServerStub adds universal methods to RootServerStubMethods.
type RootServerStub interface {
	RootServerStubMethods
	// Describe the Root interfaces.
	Describe__() []rpc.InterfaceDesc
}

// RootServer returns a server stub for Root.
// It converts an implementation of RootServerMethods into
// an object that may be used by rpc.Server.
func RootServer(impl RootServerMethods) RootServerStub {
	stub := implRootServerStub{
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

type implRootServerStub struct {
	impl RootServerMethods
	gs   *rpc.GlobState
}

func (s implRootServerStub) SeekPrivateKeys(ctx *context.T, call rpc.ServerCall) ([]bcrypter.WirePrivateKey, error) {
	return s.impl.SeekPrivateKeys(ctx, call)
}

func (s implRootServerStub) Params(ctx *context.T, call rpc.ServerCall) (bcrypter.WireParams, error) {
	return s.impl.Params(ctx, call)
}

func (s implRootServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implRootServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{RootDesc}
}

// RootDesc describes the Root interface.
var RootDesc rpc.InterfaceDesc = descRoot

// descRoot hides the desc to keep godoc clean.
var descRoot = rpc.InterfaceDesc{
	Name:    "Root",
	PkgPath: "v.io/x/ref/services/iberoot",
	Doc:     "// Root is an interface for requesting private keys for blessings.\n//\n// The keys are extracted in a blessings-based encryption scheme, which in\n// turn is based on an identity-based encryption (IBE) scheme (e.g., the BB1\n// IBE scheme).",
	Methods: []rpc.MethodDesc{
		{
			Name: "SeekPrivateKeys",
			Doc:  "// SeekPrivateKeys creates and returns private keys for blessings\n// presented by the calling principal. The blessings must be from\n// an identity provider recognized by this service.\n//\n// The extracted private keys can be used to decrypt any ciphertext\n// encrypted for a pattern matched by the presented blessings.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // []bcrypter.WirePrivateKey
			},
		},
		{
			Name: "Params",
			Doc:  "// Params returns the public encryption parameters of this service.",
			OutArgs: []rpc.ArgDesc{
				{"", ``}, // bcrypter.WireParams
			},
		},
	},
}
