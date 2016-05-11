// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: nativetest

// Package nativetest tests a package with native type conversions.
package nativetest

import (
	"fmt"
	"time"
	"v.io/v23/vdl"
	"v.io/v23/vdl/testdata/nativetest"
	"v.io/v23/vdl/vdlconv"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type WireString int32

func (WireString) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireString"`
}) {
}

func (m *WireString) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *WireString) MakeVDLTarget() vdl.Target {
	return nil
}

type WireStringTarget struct {
	Value     *string
	wireValue WireString
	vdl.TargetBase
}

func (t *WireStringTarget) FromUint(src uint64, tt *vdl.Type) error {
	t.wireValue = WireString(0)
	val, err := vdlconv.Uint64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = string(val)

	if err := WireStringToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireStringTarget) FromInt(src int64, tt *vdl.Type) error {
	t.wireValue = WireString(0)
	val, err := vdlconv.Int64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = string(val)

	if err := WireStringToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireStringTarget) FromFloat(src float64, tt *vdl.Type) error {
	t.wireValue = WireString(0)
	val, err := vdlconv.Float64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = string(val)

	if err := WireStringToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

func (x WireString) VDLIsZero() bool {
	return x == 0
}

func (x WireString) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireString)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireString) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeInt(32)
	if err != nil {
		return err
	}
	*x = WireString(tmp)
	return dec.FinishValue()
}

type WireTime int32

func (WireTime) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireTime"`
}) {
}

func (m *WireTime) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *WireTime) MakeVDLTarget() vdl.Target {
	return nil
}

type WireTimeTarget struct {
	Value     *time.Time
	wireValue WireTime
	vdl.TargetBase
}

func (t *WireTimeTarget) FromUint(src uint64, tt *vdl.Type) error {
	t.wireValue = WireTime(0)
	val, err := vdlconv.Uint64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = time.Time(val)

	if err := WireTimeToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireTimeTarget) FromInt(src int64, tt *vdl.Type) error {
	t.wireValue = WireTime(0)
	val, err := vdlconv.Int64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = time.Time(val)

	if err := WireTimeToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireTimeTarget) FromFloat(src float64, tt *vdl.Type) error {
	t.wireValue = WireTime(0)
	val, err := vdlconv.Float64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = time.Time(val)

	if err := WireTimeToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

func (x WireTime) VDLIsZero() bool {
	return x == 0
}

func (x WireTime) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireTime)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireTime) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeInt(32)
	if err != nil {
		return err
	}
	*x = WireTime(tmp)
	return dec.FinishValue()
}

type WireSamePkg int32

func (WireSamePkg) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireSamePkg"`
}) {
}

func (m *WireSamePkg) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *WireSamePkg) MakeVDLTarget() vdl.Target {
	return nil
}

type WireSamePkgTarget struct {
	Value     *nativetest.NativeSamePkg
	wireValue WireSamePkg
	vdl.TargetBase
}

func (t *WireSamePkgTarget) FromUint(src uint64, tt *vdl.Type) error {
	t.wireValue = WireSamePkg(0)
	val, err := vdlconv.Uint64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = nativetest.NativeSamePkg(val)

	if err := WireSamePkgToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireSamePkgTarget) FromInt(src int64, tt *vdl.Type) error {
	t.wireValue = WireSamePkg(0)
	val, err := vdlconv.Int64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = nativetest.NativeSamePkg(val)

	if err := WireSamePkgToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireSamePkgTarget) FromFloat(src float64, tt *vdl.Type) error {
	t.wireValue = WireSamePkg(0)
	val, err := vdlconv.Float64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = nativetest.NativeSamePkg(val)

	if err := WireSamePkgToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

func (x WireSamePkg) VDLIsZero() bool {
	return x == 0
}

func (x WireSamePkg) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireSamePkg)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireSamePkg) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeInt(32)
	if err != nil {
		return err
	}
	*x = WireSamePkg(tmp)
	return dec.FinishValue()
}

type WireMultiImport int32

func (WireMultiImport) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireMultiImport"`
}) {
}

func (m *WireMultiImport) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *WireMultiImport) MakeVDLTarget() vdl.Target {
	return nil
}

type WireMultiImportTarget struct {
	Value     *map[nativetest.NativeSamePkg]time.Time
	wireValue WireMultiImport
	vdl.TargetBase
}

func (t *WireMultiImportTarget) FromUint(src uint64, tt *vdl.Type) error {
	t.wireValue = WireMultiImport(0)
	val, err := vdlconv.Uint64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = map[nativetest.NativeSamePkg]time.Time(val)

	if err := WireMultiImportToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireMultiImportTarget) FromInt(src int64, tt *vdl.Type) error {
	t.wireValue = WireMultiImport(0)
	val, err := vdlconv.Int64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = map[nativetest.NativeSamePkg]time.Time(val)

	if err := WireMultiImportToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}
func (t *WireMultiImportTarget) FromFloat(src float64, tt *vdl.Type) error {
	t.wireValue = WireMultiImport(0)
	val, err := vdlconv.Float64ToInt32(src)
	if err != nil {
		return err
	}
	t.wireValue = map[nativetest.NativeSamePkg]time.Time(val)

	if err := WireMultiImportToNative(t.wireValue, t.Value); err != nil {
		return err
	}
	return nil
}

func (x WireMultiImport) VDLIsZero() bool {
	return x == 0
}

func (x WireMultiImport) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireMultiImport)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireMultiImport) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeInt(32)
	if err != nil {
		return err
	}
	*x = WireMultiImport(tmp)
	return dec.FinishValue()
}

type WireRenameMe int32

func (WireRenameMe) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireRenameMe"`
}) {
}

func (m *WireRenameMe) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *WireRenameMe) MakeVDLTarget() vdl.Target {
	return &WireRenameMeTarget{Value: m}
}

type WireRenameMeTarget struct {
	Value *WireRenameMe
	vdl.TargetBase
}

func (t *WireRenameMeTarget) FromUint(src uint64, tt *vdl.Type) error {

	val, err := vdlconv.Uint64ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = WireRenameMe(val)

	return nil
}
func (t *WireRenameMeTarget) FromInt(src int64, tt *vdl.Type) error {

	val, err := vdlconv.Int64ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = WireRenameMe(val)

	return nil
}
func (t *WireRenameMeTarget) FromFloat(src float64, tt *vdl.Type) error {

	val, err := vdlconv.Float64ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = WireRenameMe(val)

	return nil
}

func (x WireRenameMe) VDLIsZero() bool {
	return x == 0
}

func (x WireRenameMe) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireRenameMe)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireRenameMe) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeInt(32)
	if err != nil {
		return err
	}
	*x = WireRenameMe(tmp)
	return dec.FinishValue()
}

type WireAll struct {
	A string
	B time.Time
	C nativetest.NativeSamePkg
	D map[nativetest.NativeSamePkg]time.Time
	E WireRenameMe
}

func (WireAll) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireAll"`
}) {
}

func (m *WireAll) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}
	var wireValue2 WireString
	if err := WireStringFromNative(&wireValue2, m.A); err != nil {
		return err
	}

	var5 := (wireValue2 == WireString(0))
	if var5 {
		if err := fieldsTarget1.ZeroField("A"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("A")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue2.FillVDLTarget(fieldTarget4, tt.NonOptional().Field(0).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
				return err
			}
		}
	}
	var wireValue6 WireTime
	if err := WireTimeFromNative(&wireValue6, m.B); err != nil {
		return err
	}

	var9 := (wireValue6 == WireTime(0))
	if var9 {
		if err := fieldsTarget1.ZeroField("B"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget7, fieldTarget8, err := fieldsTarget1.StartField("B")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue6.FillVDLTarget(fieldTarget8, tt.NonOptional().Field(1).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget7, fieldTarget8); err != nil {
				return err
			}
		}
	}
	var wireValue10 WireSamePkg
	if err := WireSamePkgFromNative(&wireValue10, m.C); err != nil {
		return err
	}

	var13 := (wireValue10 == WireSamePkg(0))
	if var13 {
		if err := fieldsTarget1.ZeroField("C"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget11, fieldTarget12, err := fieldsTarget1.StartField("C")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue10.FillVDLTarget(fieldTarget12, tt.NonOptional().Field(2).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget11, fieldTarget12); err != nil {
				return err
			}
		}
	}
	var wireValue14 WireMultiImport
	if err := WireMultiImportFromNative(&wireValue14, m.D); err != nil {
		return err
	}

	var17 := (wireValue14 == WireMultiImport(0))
	if var17 {
		if err := fieldsTarget1.ZeroField("D"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("D")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := wireValue14.FillVDLTarget(fieldTarget16, tt.NonOptional().Field(3).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
				return err
			}
		}
	}
	var20 := (m.E == WireRenameMe(0))
	if var20 {
		if err := fieldsTarget1.ZeroField("E"); err != nil && err != vdl.ErrFieldNoExist {
			return err
		}
	} else {
		keyTarget18, fieldTarget19, err := fieldsTarget1.StartField("E")
		if err != vdl.ErrFieldNoExist {
			if err != nil {
				return err
			}

			if err := m.E.FillVDLTarget(fieldTarget19, tt.NonOptional().Field(4).Type); err != nil {
				return err
			}
			if err := fieldsTarget1.FinishField(keyTarget18, fieldTarget19); err != nil {
				return err
			}
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *WireAll) MakeVDLTarget() vdl.Target {
	return &WireAllTarget{Value: m}
}

type WireAllTarget struct {
	Value   *WireAll
	aTarget WireStringTarget
	bTarget WireTimeTarget
	cTarget WireSamePkgTarget
	dTarget WireMultiImportTarget
	eTarget WireRenameMeTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireAllTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {

	if ttWant := vdl.TypeOf((*WireAll)(nil)).Elem(); !vdl.Compatible(tt, ttWant) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	return t, nil
}
func (t *WireAllTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "A":
		t.aTarget.Value = &t.Value.A
		target, err := &t.aTarget, error(nil)
		return nil, target, err
	case "B":
		t.bTarget.Value = &t.Value.B
		target, err := &t.bTarget, error(nil)
		return nil, target, err
	case "C":
		t.cTarget.Value = &t.Value.C
		target, err := &t.cTarget, error(nil)
		return nil, target, err
	case "D":
		t.dTarget.Value = &t.Value.D
		target, err := &t.dTarget, error(nil)
		return nil, target, err
	case "E":
		t.eTarget.Value = &t.Value.E
		target, err := &t.eTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, vdl.ErrFieldNoExist
	}
}
func (t *WireAllTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireAllTarget) ZeroField(name string) error {
	switch name {
	case "A":
		t.Value.A = string("")
		return nil
	case "B":
		t.Value.B = time.Time{}
		return nil
	case "C":
		t.Value.C = nativetest.NativeSamePkg("")
		return nil
	case "D":
		t.Value.D = map[nativetest.NativeSamePkg]time.Time(nil)
		return nil
	case "E":
		t.Value.E = WireRenameMe(0)
		return nil
	default:
		return vdl.ErrFieldNoExist
	}
}
func (t *WireAllTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

func (x WireAll) VDLIsZero() bool {
	if x.A != "" {
		return false
	}
	if !x.B.IsZero() {
		return false
	}
	if x.C != "" {
		return false
	}
	if x.D != nil {
		return false
	}
	if x.E != 0 {
		return false
	}
	return true
}

func (x WireAll) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*WireAll)(nil)).Elem()); err != nil {
		return err
	}
	if x.A != "" {
		if err := enc.NextField("A"); err != nil {
			return err
		}
		var wire WireString
		if err := WireStringFromNative(&wire, x.A); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.B.IsZero() {
		if err := enc.NextField("B"); err != nil {
			return err
		}
		var wire WireTime
		if err := WireTimeFromNative(&wire, x.B); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.C != "" {
		if err := enc.NextField("C"); err != nil {
			return err
		}
		var wire WireSamePkg
		if err := WireSamePkgFromNative(&wire, x.C); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.D != nil {
		if err := enc.NextField("D"); err != nil {
			return err
		}
		var wire WireMultiImport
		if err := WireMultiImportFromNative(&wire, x.D); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.E != 0 {
		if err := enc.NextField("E"); err != nil {
			return err
		}
		if err := x.E.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireAll) VDLRead(dec vdl.Decoder) error {
	*x = WireAll{}
	if err := dec.StartValue(); err != nil {
		return err
	}
	if (dec.StackDepth() == 1 || dec.IsAny()) && !vdl.Compatible(vdl.TypeOf(*x), dec.Type()) {
		return fmt.Errorf("incompatible struct %T, from %v", *x, dec.Type())
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "A":
			var wire WireString
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireStringToNative(wire, &x.A); err != nil {
				return err
			}
		case "B":
			var wire WireTime
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireTimeToNative(wire, &x.B); err != nil {
				return err
			}
		case "C":
			var wire WireSamePkg
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireSamePkgToNative(wire, &x.C); err != nil {
				return err
			}
		case "D":
			var wire WireMultiImport
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireMultiImportToNative(wire, &x.D); err != nil {
				return err
			}
		case "E":
			if err := x.E.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

type ignoreme string

func (ignoreme) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.ignoreme"`
}) {
}

func (m *ignoreme) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromString(string((*m)), tt); err != nil {
		return err
	}
	return nil
}

func (m *ignoreme) MakeVDLTarget() vdl.Target {
	return &ignoremeTarget{Value: m}
}

type ignoremeTarget struct {
	Value *ignoreme
	vdl.TargetBase
}

func (t *ignoremeTarget) FromString(src string, tt *vdl.Type) error {

	if ttWant := vdl.TypeOf((*ignoreme)(nil)); !vdl.Compatible(tt, ttWant) {
		return fmt.Errorf("type %v incompatible with %v", tt, ttWant)
	}
	*t.Value = ignoreme(src)

	return nil
}

func (x ignoreme) VDLIsZero() bool {
	return x == ""
}

func (x ignoreme) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(vdl.TypeOf((*ignoreme)(nil))); err != nil {
		return err
	}
	if err := enc.EncodeString(string(x)); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *ignoreme) VDLRead(dec vdl.Decoder) error {
	if err := dec.StartValue(); err != nil {
		return err
	}
	tmp, err := dec.DecodeString()
	if err != nil {
		return err
	}
	*x = ignoreme(tmp)
	return dec.FinishValue()
}

// Type-check native conversion functions.
var (
	_ func(WireMultiImport, *map[nativetest.NativeSamePkg]time.Time) error = WireMultiImportToNative
	_ func(*WireMultiImport, map[nativetest.NativeSamePkg]time.Time) error = WireMultiImportFromNative
	_ func(WireSamePkg, *nativetest.NativeSamePkg) error                   = WireSamePkgToNative
	_ func(*WireSamePkg, nativetest.NativeSamePkg) error                   = WireSamePkgFromNative
	_ func(WireString, *string) error                                      = WireStringToNative
	_ func(*WireString, string) error                                      = WireStringFromNative
	_ func(WireTime, *time.Time) error                                     = WireTimeToNative
	_ func(*WireTime, time.Time) error                                     = WireTimeFromNative
)

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

	// Register native type conversions first, so that vdl.TypeOf works.
	vdl.RegisterNative(WireMultiImportToNative, WireMultiImportFromNative)
	vdl.RegisterNative(WireSamePkgToNative, WireSamePkgFromNative)
	vdl.RegisterNative(WireStringToNative, WireStringFromNative)
	vdl.RegisterNative(WireTimeToNative, WireTimeFromNative)

	// Register types.
	vdl.Register((*WireString)(nil))
	vdl.Register((*WireTime)(nil))
	vdl.Register((*WireSamePkg)(nil))
	vdl.Register((*WireMultiImport)(nil))
	vdl.Register((*WireRenameMe)(nil))
	vdl.Register((*WireAll)(nil))
	vdl.Register((*ignoreme)(nil))

	return struct{}{}
}
