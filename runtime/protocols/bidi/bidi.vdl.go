// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: bidi

package bidi

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Error definitions
var (
	ErrCannotListenOnBidi     = verror.Register("v.io/x/ref/runtime/protocols/bidi.CannotListenOnBidi", verror.NoRetry, "{1:}{2:} cannot listen on bidi protocol")
	ErrBidiRoutingIdNotCached = verror.Register("v.io/x/ref/runtime/protocols/bidi.BidiRoutingIdNotCached", verror.NoRetry, "{1:}{2:} bidi routing id not in cache")
)

// NewErrCannotListenOnBidi returns an error with the ErrCannotListenOnBidi ID.
func NewErrCannotListenOnBidi(ctx *context.T) error {
	return verror.New(ErrCannotListenOnBidi, ctx)
}

// NewErrBidiRoutingIdNotCached returns an error with the ErrBidiRoutingIdNotCached ID.
func NewErrBidiRoutingIdNotCached(ctx *context.T) error {
	return verror.New(ErrBidiRoutingIdNotCached, ctx)
}

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrCannotListenOnBidi.ID), "{1:}{2:} cannot listen on bidi protocol")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrBidiRoutingIdNotCached.ID), "{1:}{2:} bidi routing id not in cache")

	return struct{}{}
}