// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package manager

import (
	"bytes"
	"testing"
)

func TestFramer(t *testing.T) {
	b := &bytes.Buffer{}
	f := &framer{ReadWriter: b}
	bufs := [][]byte{[]byte("read "), []byte("this "), []byte("please.")}
	want := []byte("read this please.")
	l := len(want)
	if n, err := f.WriteMsg(bufs...); err != nil || n != l {
		t.Fatalf("got %v, %v, want %v, nil", n, err, l)
	}
	if got, err := f.ReadMsg(); err != nil || !bytes.Equal(got, want) {
		t.Errorf("got %v, %v, want %v, nil", got, err, want)
	}
	// Framing a smaller message afterwards should reuse the internal buffer
	// from the first sent message.
	bufs = [][]byte{[]byte("read "), []byte("this "), []byte("too.")}
	want = []byte("read this too.")
	oldBufferLen := l + 3
	l = len(want)
	if n, err := f.WriteMsg(bufs...); err != nil || n != l {
		t.Fatalf("got %v, %v, want %v, nil", n, err, l)
	}
	if got, err := f.ReadMsg(); err != nil || !bytes.Equal(got, want) {
		t.Errorf("got %v, %v, want %v, nil", got, err, want)
	}
	if len(f.buf) != oldBufferLen {
		t.Errorf("framer internal buffer should have been reused")
	}
	// Sending larger message afterwards should work as well.
	bufs = [][]byte{[]byte("read "), []byte("this "), []byte("way bigger message.")}
	want = []byte("read this way bigger message.")
	l = len(want)
	if n, err := f.WriteMsg(bufs...); err != nil || n != l {
		t.Fatalf("got %v, %v, want %v, nil", n, err, l)
	}
	if got, err := f.ReadMsg(); err != nil || !bytes.Equal(got, want) {
		t.Errorf("got %v, %v, want %v, nil", got, err, want)
	}
}
