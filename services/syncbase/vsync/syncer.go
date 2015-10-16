// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vsync

import (
	"time"

	"v.io/v23/context"
	"v.io/v23/verror"
	"v.io/x/lib/vlog"
)

// Policies to pick a peer to sync with.
const (
	// Picks a peer at random from the available set.
	selectRandom = iota

	// TODO(hpucha): implement other policies.
	// Picks a peer with most differing generations.
	selectMostDiff

	// Picks a peer that was synced with the furthest in the past.
	selectOldest
)

var (
	// peerSyncInterval is the duration between two consecutive peer
	// contacts. During every peer contact, the initiator obtains any
	// pending updates from that peer.
	peerSyncInterval = 50 * time.Millisecond

	// peerSelectionPolicy is the policy used to select a peer when
	// the initiator gets a chance to sync.
	peerSelectionPolicy = selectRandom

	// We wait for connectionTimeOut duration for a connection to be
	// established with a peer reachable via a chosen mount table.
	connectionTimeOut = 2 * time.Second
)

// syncer wakes up every peerSyncInterval to do work: (1) Refresh memberView if
// needed and pick a peer from all the known remote peers to sync with. (2) Act
// as an initiator and sync syncgroup metadata for all common syncgroups with
// the chosen peer (getting updates from the remote peer, detecting and
// resolving conflicts) (3) Act as an initiator and sync data corresponding to
// all common syncgroups across all Apps/Databases with the chosen peer; (4)
// Fetch any queued blobs. (5) Transfer ownership of blobs if needed. (6) Act as
// a syncgroup publisher to publish pending syncgroups; (6) Garbage collect
// older generations.
//
// TODO(hpucha): Currently only does initiation. Add rest.
func (s *syncService) syncer(ctx *context.T) {
	defer s.pending.Done()

	ticker := time.NewTicker(peerSyncInterval)
	defer ticker.Stop()

	for !s.Closed() {
		select {
		case <-ticker.C:
			if s.Closed() {
				break
			}
			s.syncerWork(ctx)

		case <-s.closed:
			break
		}
	}

	vlog.VI(1).Info("sync: syncer: channel closed, stop work and exit")
}

func (s *syncService) syncerWork(ctx *context.T) {
	// TODO(hpucha): Cut a gen for the responder even if there is no
	// one to initiate to?

	peer, err := s.pickPeer(ctx)
	if err != nil {
		return
	}

	s.syncClock(ctx, peer)

	// Sync syncgroup metadata and data.
	s.getDeltas(ctx, peer)
}

////////////////////////////////////////
// Peer selection policies.

// pickPeer picks a Syncbase to sync with.
func (s *syncService) pickPeer(ctx *context.T) (string, error) {
	switch peerSelectionPolicy {
	case selectRandom:
		members := s.getMembers(ctx)
		// Remove myself from the set.
		delete(members, s.name)
		if len(members) == 0 {
			return "", verror.New(verror.ErrInternal, ctx, "no useful peer")
		}

		// Pick a peer at random.
		ind := randIntn(len(members))
		for m := range members {
			if ind == 0 {
				return m, nil
			}
			ind--
		}
		return "", verror.New(verror.ErrInternal, ctx, "random selection didn't succeed")
	default:
		return "", verror.New(verror.ErrInternal, ctx, "unknown peer selection policy")
	}
}
