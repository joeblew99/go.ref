// This file was auto-generated by the veyron vdl tool.
// Source: otherfile.vdl

package nativetest

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"time"
	"v.io/v23/vdl/testdata/nativetest"
)

type ignoreme string

func (ignoreme) __VDLReflect(struct {
	Name string "v.io/core/veyron/lib/vdl/testdata/nativetest.ignoreme"
}) {
}

func init() {
	vdl.Register((*ignoreme)(nil))
}
