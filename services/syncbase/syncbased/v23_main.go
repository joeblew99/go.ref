// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !mojo

// The following enables go generate to generate the doc.go file.
//go:generate go run $JIRI_ROOT/release/go/src/v.io/x/lib/cmdline/testdata/gendoc.go . -help

// Package main implements syncbased, the Syncbase daemon. In addition, it
// exports MojoMain, enabling Syncbase to run as a Mojo service.
package main

import (
	"v.io/v23/context"
	"v.io/x/lib/cmdline"
	"v.io/x/ref/lib/v23cmd"
	_ "v.io/x/ref/runtime/factories/roaming"
	"v.io/x/ref/services/syncbase/syncbaselib"
)

var opts = syncbaselib.Opts{}

func main() {
	opts.InitFlags(&cmd.Flags)
	cmdline.HideGlobalFlagsExcept()
	cmdline.Main(cmd)
}

var cmd = &cmdline.Command{
	Runner: v23cmd.RunnerFunc(run),
	Name:   "syncbased",
	Short:  "Runs the Syncbase daemon",
	Long: `
Command syncbased runs the Syncbase daemon, which implements the
v.io/v23/services/syncbase interfaces.
`,
}

func run(ctx *context.T, env *cmdline.Env, args []string) error {
	syncbaselib.MainWithCtx(ctx, opts)
	return nil
}
