// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"bytes"
	"fmt"
	"sort"
	"sync"

	"v.io/v23/security"
	"v.io/v23/verror"
	"v.io/x/ref/lib/security/serialization"
)

// blessingRoots implements security.BlessingRoots.
type blessingRoots struct {
	persistedData SerializerReaderWriter
	signer        serialization.Signer
	mu            sync.RWMutex
	state         blessingRootsState // GUARDED_BY(mu)
}

func stateMapKey(root security.PublicKey) (string, error) {
	rootBytes, err := root.MarshalBinary()
	if err != nil {
		return "", err
	}
	return string(rootBytes), nil
}

func (br *blessingRoots) Add(root security.PublicKey, pattern security.BlessingPattern) error {
	key, err := stateMapKey(root)
	if err != nil {
		return err
	}

	br.mu.Lock()
	defer br.mu.Unlock()
	patterns := br.state[key]
	for _, p := range patterns {
		if p == pattern {
			return nil
		}
	}
	br.state[key] = append(patterns, pattern)

	if err := br.save(); err != nil {
		br.state[key] = patterns[:len(patterns)-1]
		return err
	}
	return nil
}

func (br *blessingRoots) Recognized(root security.PublicKey, blessing string) error {
	key, err := stateMapKey(root)
	if err != nil {
		return err
	}

	br.mu.RLock()
	defer br.mu.RUnlock()
	for _, p := range br.state[key] {
		if p.MatchedBy(blessing) {
			return nil
		}
	}
	return security.NewErrUnrecognizedRoot(nil, root.String(), nil)
}

// DebugString return a human-readable string encoding of the roots
// DebugString encodes all roots into a string in the following
// format
//
// Public key     Pattern
// <public key>   <patterns>
// ...
// <public key>   <patterns>
func (br *blessingRoots) DebugString() string {
	const format = "%-47s   %s\n"
	b := bytes.NewBufferString(fmt.Sprintf(format, "Public key", "Pattern"))
	var s rootSorter
	for keyBytes, patterns := range br.state {
		key, err := security.UnmarshalPublicKey([]byte(keyBytes))
		if err != nil {
			return fmt.Sprintf("failed to decode public key: %v", err)
		}
		s = append(s, &root{key, fmt.Sprintf("%v", patterns)})
	}
	sort.Sort(s)
	for _, r := range s {
		b.WriteString(fmt.Sprintf(format, r.key, r.patterns))
	}
	return b.String()
}

type root struct {
	key      security.PublicKey
	patterns string
}

type rootSorter []*root

func (s rootSorter) Len() int           { return len(s) }
func (s rootSorter) Less(i, j int) bool { return s[i].patterns < s[j].patterns }
func (s rootSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (br *blessingRoots) save() error {
	if (br.signer == nil) && (br.persistedData == nil) {
		return nil
	}
	data, signature, err := br.persistedData.Writers()
	if err != nil {
		return err
	}
	return encodeAndStore(br.state, data, signature, br.signer)
}

// newInMemoryBlessingRoots returns an in-memory security.BlessingRoots.
//
// The returned BlessingRoots is initialized with an empty set of keys.
func newInMemoryBlessingRoots() security.BlessingRoots {
	return &blessingRoots{
		state: make(blessingRootsState),
	}
}

// newPersistingBlessingRoots returns a security.BlessingRoots for a principal
// that is initialized with the persisted data. The returned security.BlessingRoots
// also persists any updates to its state.
func newPersistingBlessingRoots(persistedData SerializerReaderWriter, signer serialization.Signer) (security.BlessingRoots, error) {
	if persistedData == nil || signer == nil {
		return nil, verror.New(errDataOrSignerUnspecified, nil)
	}
	br := &blessingRoots{
		state:         make(blessingRootsState),
		persistedData: persistedData,
		signer:        signer,
	}
	data, signature, err := br.persistedData.Readers()
	if err != nil {
		return nil, err
	}
	if (data != nil) && (signature != nil) {
		if err := decodeFromStorage(&br.state, data, signature, br.signer.PublicKey()); err != nil {
			return nil, err
		}
	}
	return br, nil
}
