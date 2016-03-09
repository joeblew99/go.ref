// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: errors.vdl

package internal

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/security"
	"v.io/v23/verror"
)

func __VDLEnsureNativeBuilt_errors() {
}

var (
	ErrUnrecognizedRemoteBlessings = verror.Register("v.io/x/ref/services/iberoot/rootd/internal.UnrecognizedRemoteBlessings", verror.NoRetry, "{1:}{2:} blessing provided by the remote end: {3} [rejected: {4}] are not recognized by this identity provider: {5}")
	ErrInternal                    = verror.Register("v.io/x/ref/services/iberoot/rootd/internal.Internal", verror.NoRetry, "{1:}{2:} internal error: {3}")
)

func init() {
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrUnrecognizedRemoteBlessings.ID), "{1:}{2:} blessing provided by the remote end: {3} [rejected: {4}] are not recognized by this identity provider: {5}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInternal.ID), "{1:}{2:} internal error: {3}")
}

// NewErrUnrecognizedRemoteBlessings returns an error with the ErrUnrecognizedRemoteBlessings ID.
func NewErrUnrecognizedRemoteBlessings(ctx *context.T, blessings []string, rejected []security.RejectedBlessing, name string) error {
	return verror.New(ErrUnrecognizedRemoteBlessings, ctx, blessings, rejected, name)
}

// NewErrInternal returns an error with the ErrInternal ID.
func NewErrInternal(ctx *context.T, err error) error {
	return verror.New(ErrInternal, ctx, err)
}
