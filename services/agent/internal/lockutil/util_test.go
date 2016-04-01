// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lockutil_test contains tests for the lockutil package.
package lockutil_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"v.io/x/lib/gosh"
	"v.io/x/ref/services/agent/internal/lockutil"
)

// TestCreateLockFile ensures that CreateLockFile writes the file it's supposed
// to.
func TestCreateLockFile(t *testing.T) {
	d, err := ioutil.TempDir("", "lockutiltest")
	if err != nil {
		t.Fatalf("TempDir failed: %v", err)
	}
	defer os.RemoveAll(d)
	f, err := lockutil.CreateLockFile(d, "foo")
	if err != nil {
		t.Fatalf("createPIDFile failed: %v", err)
	}
	files, err := ioutil.ReadDir(d)
	if err != nil {
		t.Fatalf("ReadDir failed: %v", err)
	}
	if nfiles := len(files); nfiles != 1 {
		t.Fatalf("Expected 1 file, found %d", nfiles)
	}
	if found, want := files[0].Name(), filepath.Base(f); found != want {
		t.Fatalf("Expected file %s, found %s instead", want, found)
	}
}

var goshCreateLockFile = gosh.RegisterFunc("CreateLockFile", func(dir string) error {
	f, err := lockutil.CreateLockFile(dir, "foo")
	if err != nil {
		return err
	}
	fmt.Println(f)
	return nil
})

// TestStillHeld verifies StillHeld returns the appropriate boolean when
// presented with either a running or a dead process' information.
func TestStillHeld(t *testing.T) {
	d, err := ioutil.TempDir("", "lockutiltest")
	if err != nil {
		t.Fatalf("TempDir failed: %v", err)
	}
	defer os.RemoveAll(d)

	f, err := lockutil.CreateLockFile(d, "foo")
	if err != nil {
		t.Fatalf("createPIDFile failed: %v", err)
	}
	if info, err := ioutil.ReadFile(f); err != nil {
		t.Fatalf("ReadFile(%v) failed: %v", f, err)
	} else if running, err := lockutil.StillHeld(info); err != nil || !running {
		t.Fatalf("Expected (true, <nil>) got (%t, %v) instead from StillHeld for:\n%v", running, err, string(info))
	}

	sh := gosh.NewShell(t)
	defer sh.Cleanup()
	if out := sh.FuncCmd(goshCreateLockFile, d).Stdout(); filepath.Dir(out) != d {
		t.Fatalf("Unexpected output: %s", out)
	} else {
		f = strings.TrimSuffix(out, "\n")
	}
	if info, err := ioutil.ReadFile(f); err != nil {
		t.Fatalf("ReadFile(%v) failed: %v", f, err)
	} else if running, err := lockutil.StillHeld(info); err != nil || running {
		t.Fatalf("Expected (false, <nil>) got (%t, %v) instead from StillHeld for:\n%v", running, err, string(info))
	}
}

func TestInvalidVersion(t *testing.T) {
	for i, c := range []struct {
		info, error string
	}{
		{
			"VERSION:1000000\nBLAH\n",
			"unknown version: 1000000",
		},
		{
			"VERSION:1\nBLAH\n",
			"failed to parse",
		},
	} {
		if _, err := lockutil.StillHeld([]byte(c.info)); err == nil {
			t.Fatalf("case %d: expected to fail\n", i)
		} else if !strings.Contains(err.Error(), c.error) {
			t.Fatalf("case %d: expected error \"...%s...\", got \"%s\" instead", i, c.error, err.Error())
		}
	}
}

func TestMain(m *testing.M) {
	gosh.InitMain()
	os.Exit(m.Run())
}
