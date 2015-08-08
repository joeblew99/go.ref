// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conn

import (
	"bytes"
	"fmt"
	"testing"

	"v.io/v23"
	_ "v.io/x/ref/runtime/factories/fake"
)

func TestRemoteDialerClose(t *testing.T) {
	ctx, shutdown := v23.Init()
	defer shutdown()
	d, a, w := setupConns(t, ctx, nil, nil)
	d.Close(ctx, fmt.Errorf("Closing randomly."))
	<-d.Closed()
	<-a.Closed()
	if !w.isClosed() {
		t.Errorf("The connection should be closed")
	}
}

func TestRemoteAcceptorClose(t *testing.T) {
	ctx, shutdown := v23.Init()
	defer shutdown()
	d, a, w := setupConns(t, ctx, nil, nil)
	a.Close(ctx, fmt.Errorf("Closing randomly."))
	<-a.Closed()
	<-d.Closed()
	if !w.isClosed() {
		t.Errorf("The connection should be closed")
	}
}

func TestUnderlyingConnectionClosed(t *testing.T) {
	ctx, shutdown := v23.Init()
	defer shutdown()
	d, a, w := setupConns(t, ctx, nil, nil)
	w.close()
	<-a.Closed()
	<-d.Closed()
}

func TestDialAfterConnClose(t *testing.T) {
	ctx, shutdown := v23.Init()
	defer shutdown()
	d, a, _ := setupConns(t, ctx, nil, nil)

	d.Close(ctx, fmt.Errorf("Closing randomly."))
	<-d.Closed()
	<-a.Closed()
	if _, err := d.Dial(ctx); err == nil {
		t.Errorf("Nil error dialing on dialer")
	}
	if _, err := a.Dial(ctx); err == nil {
		t.Errorf("Nil error dialing on acceptor")
	}
}

func TestReadWriteAfterConnClose(t *testing.T) {
	ctx, shutdown := v23.Init()
	defer shutdown()
	for _, dialerDials := range []bool{true, false} {
		df, flows := setupFlow(t, ctx, dialerDials)
		if _, err := df.WriteMsg([]byte("hello")); err != nil {
			t.Fatalf("write failed: %v", err)
		}
		af := <-flows
		if got, err := af.ReadMsg(); err != nil {
			t.Fatalf("read failed: %v", err)
		} else if !bytes.Equal(got, []byte("hello")) {
			t.Errorf("got %s want %s", string(got), "hello")
		}
		if _, err := df.WriteMsg([]byte("there")); err != nil {
			t.Fatalf("second write failed: %v", err)
		}
		df.(*flw).conn.Close(ctx, nil)
		<-af.Conn().Closed()
		if got, err := af.ReadMsg(); err != nil {
			t.Fatalf("read failed: %v", err)
		} else if !bytes.Equal(got, []byte("there")) {
			t.Errorf("got %s want %s", string(got), "there")
		}
		if _, err := df.WriteMsg([]byte("fail")); err == nil {
			t.Errorf("nil error for write after close.")
		}
		if _, err := af.ReadMsg(); err == nil {
			t.Fatalf("nil error for read after close.")
		}
	}
}
