// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: types.vdl

package bcrypter

import (
	// VDL system imports
	"v.io/v23/vdl"
)

// WireCiphertext represents the wire format of the ciphertext
// generated by a Crypter.
type WireCiphertext struct {
	// PatternId is an identifier of the blessing pattern that this
	// ciphertext is for. It is represented by a 16 byte truncated
	// SHA256 hash of the pattern.
	PatternId string
	// Bytes is a map from an identifier of the public IBE params to
	// the ciphertext bytes that were generated using those params.
	//
	// The params identifier is a 16 byte truncated SHA256 hash
	// of the marshaled form of the IBE params.
	Bytes map[string][]byte
}

func (WireCiphertext) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WireCiphertext"`
}) {
}

// WireParams represents the wire format of the public parameters
// of an identity provider (aka Root).
type WireParams struct {
	// Blessing is the blessing name of the identity provider. The identity
	// provider  can extract private keys for blessings that are extensions
	// of this blessing name.
	Blessing string
	// Params is the marshaled form of the public IBE params of the
	// the identity provider.
	Params []byte
}

func (WireParams) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WireParams"`
}) {
}

// WirePrivateKey represents the wire format of the private key corresponding
// to a blessing.
type WirePrivateKey struct {
	// Blessing is the blessing for which this private key was extracted for.
	Blessing string
	// Params are the public parameters of the identity provider that extracted
	// this private key.
	Params WireParams
	// Keys contain the extracted IBE private keys for each pattern that is
	// matched by the blessing and is an extension of the identity provider's
	// name. The keys are enumerated in increasing order of the lengths of the
	// corresponding patterns.
	//
	// For example, if the blessing is "google/u/alice/phone" and the identity
	// provider's name is "google/u" then the keys are extracted for the patterns
	// - "google/u"
	// - "google/u/alice"
	// - "google/u/alice/phone"
	// - "google/u/alice/phone/$"
	//
	// The private keys are listed in increasing order of the lengths of the
	// corresponding patterns.
	Keys [][]byte
}

func (WirePrivateKey) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/security/bcrypter.WirePrivateKey"`
}) {
}

func init() {
	vdl.Register((*WireCiphertext)(nil))
	vdl.Register((*WireParams)(nil))
	vdl.Register((*WirePrivateKey)(nil))
}
