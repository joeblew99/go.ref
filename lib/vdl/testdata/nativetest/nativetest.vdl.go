// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Source: nativetest.vdl

// Package nativetest tests a package with native type conversions.
package nativetest

import (
	"fmt"
	"reflect"
	"time"
	"v.io/v23/vdl"
	"v.io/v23/vdl/testdata/nativetest"
	"v.io/v23/vdl/vdlconv"
)

type WireString int32

func (WireString) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireString"`
}) {
}

func (m *WireString) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString); err != nil {
		return err
	}
	return nil
}

func (m *WireString) MakeVDLTarget() vdl.Target {
	return nil
}

type WireMapStringInt int32

func (WireMapStringInt) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireMapStringInt"`
}) {
}

func (m *WireMapStringInt) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt); err != nil {
		return err
	}
	return nil
}

func (m *WireMapStringInt) MakeVDLTarget() vdl.Target {
	return nil
}

type WireTime int32

func (WireTime) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireTime"`
}) {
}

func (m *WireTime) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime); err != nil {
		return err
	}
	return nil
}

func (m *WireTime) MakeVDLTarget() vdl.Target {
	return nil
}

type WireSamePkg int32

func (WireSamePkg) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireSamePkg"`
}) {
}

func (m *WireSamePkg) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg); err != nil {
		return err
	}
	return nil
}

func (m *WireSamePkg) MakeVDLTarget() vdl.Target {
	return nil
}

type WireMultiImport int32

func (WireMultiImport) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireMultiImport"`
}) {
}

func (m *WireMultiImport) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport); err != nil {
		return err
	}
	return nil
}

func (m *WireMultiImport) MakeVDLTarget() vdl.Target {
	return nil
}

type WireRenameMe int32

func (WireRenameMe) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireRenameMe"`
}) {
}

func (m *WireRenameMe) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if err := t.FromInt(int64((*m)), __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireRenameMe); err != nil {
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
func (t *WireRenameMeTarget) FromComplex(src complex128, tt *vdl.Type) error {
	val, err := vdlconv.Complex128ToInt32(src)
	if err != nil {
		return err
	}
	*t.Value = WireRenameMe(val)
	return nil
}

type WireAll struct {
	A string
	B map[string]int
	C time.Time
	D nativetest.NativeSamePkg
	E map[nativetest.NativeSamePkg]time.Time
	F WireRenameMe
}

func (WireAll) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireAll"`
}) {
}

func (m *WireAll) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	__VDLEnsureNativeBuilt_nativetest()
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	var wireValue2 WireString
	if err := WireStringFromNative(&wireValue2, m.A); err != nil {
		return err
	}

	keyTarget3, fieldTarget4, err := fieldsTarget1.StartField("A")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue2.FillVDLTarget(fieldTarget4, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget3, fieldTarget4); err != nil {
			return err
		}
	}
	var wireValue5 WireMapStringInt
	if err := WireMapStringIntFromNative(&wireValue5, m.B); err != nil {
		return err
	}

	keyTarget6, fieldTarget7, err := fieldsTarget1.StartField("B")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue5.FillVDLTarget(fieldTarget7, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget6, fieldTarget7); err != nil {
			return err
		}
	}
	var wireValue8 WireTime
	if err := WireTimeFromNative(&wireValue8, m.C); err != nil {
		return err
	}

	keyTarget9, fieldTarget10, err := fieldsTarget1.StartField("C")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue8.FillVDLTarget(fieldTarget10, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget9, fieldTarget10); err != nil {
			return err
		}
	}
	var wireValue11 WireSamePkg
	if err := WireSamePkgFromNative(&wireValue11, m.D); err != nil {
		return err
	}

	keyTarget12, fieldTarget13, err := fieldsTarget1.StartField("D")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue11.FillVDLTarget(fieldTarget13, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget12, fieldTarget13); err != nil {
			return err
		}
	}
	var wireValue14 WireMultiImport
	if err := WireMultiImportFromNative(&wireValue14, m.E); err != nil {
		return err
	}

	keyTarget15, fieldTarget16, err := fieldsTarget1.StartField("E")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := wireValue14.FillVDLTarget(fieldTarget16, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget15, fieldTarget16); err != nil {
			return err
		}
	}
	keyTarget17, fieldTarget18, err := fieldsTarget1.StartField("F")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.F.FillVDLTarget(fieldTarget18, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireRenameMe); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget17, fieldTarget18); err != nil {
			return err
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
	Value *WireAll
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *WireAllTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll)
	}
	return t, nil
}
func (t *WireAllTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "A":
		val, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.A))
		return nil, val, err
	case "B":
		val, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.B))
		return nil, val, err
	case "C":
		val, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.C))
		return nil, val, err
	case "D":
		val, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.D))
		return nil, val, err
	case "E":
		val, err := vdl.ReflectTarget(reflect.ValueOf(&t.Value.E))
		return nil, val, err
	case "F":
		val, err := &WireRenameMeTarget{Value: &t.Value.F}, error(nil)
		return nil, val, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll)
	}
}
func (t *WireAllTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *WireAllTarget) FinishFields(_ vdl.FieldsTarget) error {
	return nil
}

func init() {
	vdl.RegisterNative(WireMapStringIntToNative, WireMapStringIntFromNative)
	vdl.RegisterNative(WireMultiImportToNative, WireMultiImportFromNative)
	vdl.RegisterNative(WireSamePkgToNative, WireSamePkgFromNative)
	vdl.RegisterNative(WireStringToNative, WireStringFromNative)
	vdl.RegisterNative(WireTimeToNative, WireTimeFromNative)
	vdl.Register((*WireString)(nil))
	vdl.Register((*WireMapStringInt)(nil))
	vdl.Register((*WireTime)(nil))
	vdl.Register((*WireSamePkg)(nil))
	vdl.Register((*WireMultiImport)(nil))
	vdl.Register((*WireRenameMe)(nil))
	vdl.Register((*WireAll)(nil))
}

// Type-check WireMapStringInt conversion functions.
var _ func(WireMapStringInt, *map[string]int) error = WireMapStringIntToNative
var _ func(*WireMapStringInt, map[string]int) error = WireMapStringIntFromNative

// Type-check WireMultiImport conversion functions.
var _ func(WireMultiImport, *map[nativetest.NativeSamePkg]time.Time) error = WireMultiImportToNative
var _ func(*WireMultiImport, map[nativetest.NativeSamePkg]time.Time) error = WireMultiImportFromNative

// Type-check WireSamePkg conversion functions.
var _ func(WireSamePkg, *nativetest.NativeSamePkg) error = WireSamePkgToNative
var _ func(*WireSamePkg, nativetest.NativeSamePkg) error = WireSamePkgFromNative

// Type-check WireString conversion functions.
var _ func(WireString, *string) error = WireStringToNative
var _ func(*WireString, string) error = WireStringFromNative

// Type-check WireTime conversion functions.
var _ func(WireTime, *time.Time) error = WireTimeToNative
var _ func(*WireTime, time.Time) error = WireTimeFromNative

var __VDLTypenativetest0 *vdl.Type

func __VDLTypenativetest0_gen() *vdl.Type {
	__VDLTypenativetest0Builder := vdl.TypeBuilder{}

	__VDLTypenativetest01 := __VDLTypenativetest0Builder.Optional()
	__VDLTypenativetest02 := __VDLTypenativetest0Builder.Struct()
	__VDLTypenativetest03 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireAll").AssignBase(__VDLTypenativetest02)
	__VDLTypenativetest04 := vdl.Int32Type
	__VDLTypenativetest05 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireString").AssignBase(__VDLTypenativetest04)
	__VDLTypenativetest02.AppendField("A", __VDLTypenativetest05)
	__VDLTypenativetest06 := vdl.Int32Type
	__VDLTypenativetest07 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireMapStringInt").AssignBase(__VDLTypenativetest06)
	__VDLTypenativetest02.AppendField("B", __VDLTypenativetest07)
	__VDLTypenativetest08 := vdl.Int32Type
	__VDLTypenativetest09 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireTime").AssignBase(__VDLTypenativetest08)
	__VDLTypenativetest02.AppendField("C", __VDLTypenativetest09)
	__VDLTypenativetest010 := vdl.Int32Type
	__VDLTypenativetest011 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireSamePkg").AssignBase(__VDLTypenativetest010)
	__VDLTypenativetest02.AppendField("D", __VDLTypenativetest011)
	__VDLTypenativetest012 := vdl.Int32Type
	__VDLTypenativetest013 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireMultiImport").AssignBase(__VDLTypenativetest012)
	__VDLTypenativetest02.AppendField("E", __VDLTypenativetest013)
	__VDLTypenativetest014 := vdl.Int32Type
	__VDLTypenativetest015 := __VDLTypenativetest0Builder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireRenameMe").AssignBase(__VDLTypenativetest014)
	__VDLTypenativetest02.AppendField("F", __VDLTypenativetest015)
	__VDLTypenativetest01.AssignElem(__VDLTypenativetest03)
	__VDLTypenativetest0Builder.Build()
	__VDLTypenativetest0v, err := __VDLTypenativetest01.Built()
	if err != nil {
		panic(err)
	}
	return __VDLTypenativetest0v
}
func init() {
	__VDLTypenativetest0 = __VDLTypenativetest0_gen()
}

var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll *vdl.Type

func __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll_gen() *vdl.Type {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder := vdl.TypeBuilder{}

	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Struct()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll2 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireAll").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll3 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll4 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireString").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll3)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1.AppendField("A", __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll4)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll5 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll6 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireMapStringInt").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll5)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1.AppendField("B", __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll6)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll7 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll8 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireTime").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll7)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1.AppendField("C", __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll8)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll9 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll10 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireSamePkg").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll9)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1.AppendField("D", __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll10)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll11 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll12 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireMultiImport").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll11)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1.AppendField("E", __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll12)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll13 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll14 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireRenameMe").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll13)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll1.AppendField("F", __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll14)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllBuilder.Build()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllv, err := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAllv
}
func init() {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll_gen()
}

var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt *vdl.Type

func __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt_gen() *vdl.Type {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringIntBuilder := vdl.TypeBuilder{}

	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt1 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt2 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringIntBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireMapStringInt").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt1)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringIntBuilder.Build()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringIntv, err := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringIntv
}
func init() {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt_gen()
}

var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport *vdl.Type

func __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport_gen() *vdl.Type {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImportBuilder := vdl.TypeBuilder{}

	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport1 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport2 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImportBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireMultiImport").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport1)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImportBuilder.Build()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImportv, err := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImportv
}
func init() {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport_gen()
}

var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireRenameMe *vdl.Type = vdl.TypeOf(WireRenameMe(0))
var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg *vdl.Type

func __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg_gen() *vdl.Type {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkgBuilder := vdl.TypeBuilder{}

	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg1 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg2 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkgBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireSamePkg").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg1)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkgBuilder.Build()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkgv, err := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkgv
}
func init() {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg_gen()
}

var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString *vdl.Type

func __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString_gen() *vdl.Type {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireStringBuilder := vdl.TypeBuilder{}

	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString1 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString2 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireStringBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireString").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString1)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireStringBuilder.Build()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireStringv, err := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireStringv
}
func init() {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString_gen()
}

var __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime *vdl.Type

func __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime_gen() *vdl.Type {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTimeBuilder := vdl.TypeBuilder{}

	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime1 := vdl.Int32Type
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime2 := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTimeBuilder.Named("v.io/x/ref/lib/vdl/testdata/nativetest.WireTime").AssignBase(__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime1)
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTimeBuilder.Build()
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTimev, err := __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime2.Built()
	if err != nil {
		panic(err)
	}
	return __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTimev
}
func init() {
	__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime_gen()
}
func __VDLEnsureNativeBuilt_nativetest() {
	if __VDLTypenativetest0 == nil {
		__VDLTypenativetest0 = __VDLTypenativetest0_gen()
	}
	if __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll == nil {
		__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireAll_gen()
	}
	if __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt == nil {
		__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMapStringInt_gen()
	}
	if __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport == nil {
		__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireMultiImport_gen()
	}
	if __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg == nil {
		__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireSamePkg_gen()
	}
	if __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString == nil {
		__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireString_gen()
	}
	if __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime == nil {
		__VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime = __VDLType_nativetest_v_io_x_ref_lib_vdl_testdata_nativetest_WireTime_gen()
	}
}
