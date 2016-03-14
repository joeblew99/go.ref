// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: blobtestsvdl

package blobtestsvdl

import (
	"fmt"
	"reflect"
	"v.io/v23/services/syncbase/nosql"
	"v.io/v23/vdl"
	"v.io/v23/vom"
)

type BlobInfo struct {
	Info string
	Br   nosql.BlobRef
}

func (BlobInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobInfo"`
}) {
}

func (m *BlobInfo) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Br")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Br.FillVDLTarget(fieldTarget5, __VDLType_v_io_v23_services_syncbase_nosql_BlobRef); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobInfo) MakeVDLTarget() vdl.Target {
	return &BlobInfoTarget{Value: m}
}

type BlobInfoTarget struct {
	Value      *BlobInfo
	infoTarget vdl.StringTarget
	brTarget   nosql.BlobRefTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobInfoTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo)
	}
	return t, nil
}
func (t *BlobInfoTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Br":
		t.brTarget.Value = &t.Value.Br
		target, err := &t.brTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo)
	}
}
func (t *BlobInfoTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobInfoTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

type (
	// BlobUnion represents any single field of the BlobUnion union type.
	BlobUnion interface {
		// Index returns the field index.
		Index() int
		// Interface returns the field value as an interface.
		Interface() interface{}
		// Name returns the field name.
		Name() string
		// __VDLReflect describes the BlobUnion union type.
		__VDLReflect(__BlobUnionReflect)
		FillVDLTarget(vdl.Target, *vdl.Type) error
	}
	// BlobUnionNum represents field Num of the BlobUnion union type.
	BlobUnionNum struct{ Value int32 }
	// BlobUnionBi represents field Bi of the BlobUnion union type.
	BlobUnionBi struct{ Value BlobInfo }
	// __BlobUnionReflect describes the BlobUnion union type.
	__BlobUnionReflect struct {
		Name  string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobUnion"`
		Type  BlobUnion
		Union struct {
			Num BlobUnionNum
			Bi  BlobUnionBi
		}
	}
)

func (x BlobUnionNum) Index() int                      { return 0 }
func (x BlobUnionNum) Interface() interface{}          { return x.Value }
func (x BlobUnionNum) Name() string                    { return "Num" }
func (x BlobUnionNum) __VDLReflect(__BlobUnionReflect) {}

func (m BlobUnionNum) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(__VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobUnion)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Num")
	if err != nil {
		return err
	}
	if err := fieldTarget3.FromInt(int64(m.Value), vdl.Int32Type); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlobUnionNum) MakeVDLTarget() vdl.Target {
	return nil
}

func (x BlobUnionBi) Index() int                      { return 1 }
func (x BlobUnionBi) Interface() interface{}          { return x.Value }
func (x BlobUnionBi) Name() string                    { return "Bi" }
func (x BlobUnionBi) __VDLReflect(__BlobUnionReflect) {}

func (m BlobUnionBi) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	fieldsTarget1, err := t.StartFields(__VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobUnion)
	if err != nil {
		return err
	}
	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Bi")
	if err != nil {
		return err
	}

	if err := m.Value.FillVDLTarget(fieldTarget3, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo); err != nil {
		return err
	}
	if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
		return err
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}

	return nil
}

func (m BlobUnionBi) MakeVDLTarget() vdl.Target {
	return nil
}

type BlobSet struct {
	Info string
	Bs   map[nosql.BlobRef]struct{}
}

func (BlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobSet"`
}) {
}

func (m *BlobSet) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobSet == nil || __VDLType1 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Bs")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		setTarget6, err := fieldTarget5.StartSet(__VDLType2, len(m.Bs))
		if err != nil {
			return err
		}
		for key8 := range m.Bs {
			keyTarget7, err := setTarget6.StartKey()
			if err != nil {
				return err
			}

			if err := key8.FillVDLTarget(keyTarget7, __VDLType_v_io_v23_services_syncbase_nosql_BlobRef); err != nil {
				return err
			}
			if err := setTarget6.FinishKey(keyTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishSet(setTarget6); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobSet) MakeVDLTarget() vdl.Target {
	return &BlobSetTarget{Value: m}
}

type BlobSetTarget struct {
	Value      *BlobSet
	infoTarget vdl.StringTarget
	bsTarget   unnamed_7365745b762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e675dTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobSetTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobSet) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobSet)
	}
	return t, nil
}
func (t *BlobSetTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Bs":
		t.bsTarget.Value = &t.Value.Bs
		target, err := &t.bsTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobSet)
	}
}
func (t *BlobSetTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobSetTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[nosql.BlobRef]struct{}
type unnamed_7365745b762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e675dTarget struct {
	Value     *map[nosql.BlobRef]struct{}
	currKey   nosql.BlobRef
	keyTarget nosql.BlobRefTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e675dTarget) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {
	if !vdl.Compatible(tt, __VDLType2) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType2)
	}
	*t.Value = make(map[nosql.BlobRef]struct{})
	return t, nil
}
func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e675dTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = nosql.BlobRef("")
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e675dTarget) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e675dTarget) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

type BlobAny struct {
	Info string
	Baa  []*vom.RawBytes
}

func (BlobAny) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobAny"`
}) {
}

func (m *BlobAny) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobAny == nil || __VDLType3 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Baa")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		listTarget6, err := fieldTarget5.StartList(__VDLType4, len(m.Baa))
		if err != nil {
			return err
		}
		for i, elem8 := range m.Baa {
			elemTarget7, err := listTarget6.StartElem(i)
			if err != nil {
				return err
			}

			if elem8 == nil {
				if err := elemTarget7.FromNil(vdl.AnyType); err != nil {
					return err
				}
			} else {
				if err := elem8.FillVDLTarget(elemTarget7, vdl.AnyType); err != nil {
					return err
				}
			}
			if err := listTarget6.FinishElem(elemTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishList(listTarget6); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobAny) MakeVDLTarget() vdl.Target {
	return &BlobAnyTarget{Value: m}
}

type BlobAnyTarget struct {
	Value      *BlobAny
	infoTarget vdl.StringTarget
	baaTarget  unnamed_5b5d616e79Target
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobAnyTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobAny) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobAny)
	}
	return t, nil
}
func (t *BlobAnyTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Baa":
		t.baaTarget.Value = &t.Value.Baa
		target, err := &t.baaTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobAny)
	}
}
func (t *BlobAnyTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobAnyTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// []*vom.RawBytes
type unnamed_5b5d616e79Target struct {
	Value *[]*vom.RawBytes

	vdl.TargetBase
	vdl.ListTargetBase
}

func (t *unnamed_5b5d616e79Target) StartList(tt *vdl.Type, len int) (vdl.ListTarget, error) {
	if !vdl.Compatible(tt, __VDLType4) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType4)
	}
	if cap(*t.Value) < len {
		*t.Value = make([]*vom.RawBytes, len)
	} else {
		*t.Value = (*t.Value)[:len]
	}
	return t, nil
}
func (t *unnamed_5b5d616e79Target) StartElem(index int) (elem vdl.Target, _ error) {
	target, err := vdl.ReflectTarget(reflect.ValueOf(&(*t.Value)[index]))
	return target, err
}
func (t *unnamed_5b5d616e79Target) FinishElem(elem vdl.Target) error {
	return nil
}
func (t *unnamed_5b5d616e79Target) FinishList(elem vdl.ListTarget) error {

	return nil
}

type NonBlobSet struct {
	Info string
	S    map[string]struct{}
}

func (NonBlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.NonBlobSet"`
}) {
}

func (m *NonBlobSet) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_NonBlobSet == nil || __VDLType5 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("S")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		setTarget6, err := fieldTarget5.StartSet(__VDLType6, len(m.S))
		if err != nil {
			return err
		}
		for key8 := range m.S {
			keyTarget7, err := setTarget6.StartKey()
			if err != nil {
				return err
			}
			if err := keyTarget7.FromString(string(key8), vdl.StringType); err != nil {
				return err
			}
			if err := setTarget6.FinishKey(keyTarget7); err != nil {
				return err
			}
		}
		if err := fieldTarget5.FinishSet(setTarget6); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *NonBlobSet) MakeVDLTarget() vdl.Target {
	return &NonBlobSetTarget{Value: m}
}

type NonBlobSetTarget struct {
	Value      *NonBlobSet
	infoTarget vdl.StringTarget
	sTarget    unnamed_7365745b737472696e675dTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *NonBlobSetTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_NonBlobSet) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_NonBlobSet)
	}
	return t, nil
}
func (t *NonBlobSetTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "S":
		t.sTarget.Value = &t.Value.S
		target, err := &t.sTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_NonBlobSet)
	}
}
func (t *NonBlobSetTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *NonBlobSetTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[string]struct{}
type unnamed_7365745b737472696e675dTarget struct {
	Value     *map[string]struct{}
	currKey   string
	keyTarget vdl.StringTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *unnamed_7365745b737472696e675dTarget) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {
	if !vdl.Compatible(tt, __VDLType6) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType6)
	}
	*t.Value = make(map[string]struct{})
	return t, nil
}
func (t *unnamed_7365745b737472696e675dTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = ""
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *unnamed_7365745b737472696e675dTarget) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *unnamed_7365745b737472696e675dTarget) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

type BlobOpt struct {
	Info string
	Bo   *BlobInfo
}

func (BlobOpt) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobOpt"`
}) {
}

func (m *BlobOpt) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobOpt == nil || __VDLType7 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Info")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {
		if err := fieldTarget3.FromString(string(m.Info), vdl.StringType); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Bo")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if m.Bo == nil {
			if err := fieldTarget5.FromNil(__VDLType0); err != nil {
				return err
			}
		} else {
			if err := m.Bo.FillVDLTarget(fieldTarget5, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo); err != nil {
				return err
			}
		}
		if err := fieldsTarget1.FinishField(keyTarget4, fieldTarget5); err != nil {
			return err
		}
	}
	if err := t.FinishFields(fieldsTarget1); err != nil {
		return err
	}
	return nil
}

func (m *BlobOpt) MakeVDLTarget() vdl.Target {
	return &BlobOptTarget{Value: m}
}

type BlobOptTarget struct {
	Value      *BlobOpt
	infoTarget vdl.StringTarget
	boTarget   unnamed_3f762e696f2f782f7265662f73657276696365732f73796e63626173652f7673796e632f74657374646174612e426c6f62496e666f207374727563747b496e666f20737472696e673b427220762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e677dTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *BlobOptTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobOpt) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobOpt)
	}
	return t, nil
}
func (t *BlobOptTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Info":
		t.infoTarget.Value = &t.Value.Info
		target, err := &t.infoTarget, error(nil)
		return nil, target, err
	case "Bo":
		t.boTarget.Value = &t.Value.Bo
		target, err := &t.boTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobOpt)
	}
}
func (t *BlobOptTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *BlobOptTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// Optional BlobInfo
type unnamed_3f762e696f2f782f7265662f73657276696365732f73796e63626173652f7673796e632f74657374646174612e426c6f62496e666f207374727563747b496e666f20737472696e673b427220762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e677dTarget struct {
	Value      **BlobInfo
	elemTarget BlobInfoTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *unnamed_3f762e696f2f782f7265662f73657276696365732f73796e63626173652f7673796e632f74657374646174612e426c6f62496e666f207374727563747b496e666f20737472696e673b427220762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e677dTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if *t.Value == nil {
		*t.Value = &BlobInfo{}
	}
	t.elemTarget.Value = *t.Value
	target, err := &t.elemTarget, error(nil)
	if err != nil {
		return nil, err
	}
	return target.StartFields(tt)
}
func (t *unnamed_3f762e696f2f782f7265662f73657276696365732f73796e63626173652f7673796e632f74657374646174612e426c6f62496e666f207374727563747b496e666f20737472696e673b427220762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e677dTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}
func (t *unnamed_3f762e696f2f782f7265662f73657276696365732f73796e63626173652f7673796e632f74657374646174612e426c6f62496e666f207374727563747b496e666f20737472696e673b427220762e696f2f7632332f73657276696365732f73796e63626173652f6e6f73716c2e426c6f6252656620737472696e677dTarget) FromNil(tt *vdl.Type) error {
	*t.Value = nil

	return nil
}

func init() {
	vdl.Register((*BlobInfo)(nil))
	vdl.Register((*BlobUnion)(nil))
	vdl.Register((*BlobSet)(nil))
	vdl.Register((*BlobAny)(nil))
	vdl.Register((*NonBlobSet)(nil))
	vdl.Register((*BlobOpt)(nil))
}

var __VDLType3 *vdl.Type = vdl.TypeOf((*BlobAny)(nil))
var __VDLType0 *vdl.Type = vdl.TypeOf((*BlobInfo)(nil))
var __VDLType7 *vdl.Type = vdl.TypeOf((*BlobOpt)(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf((*BlobSet)(nil))
var __VDLType5 *vdl.Type = vdl.TypeOf((*NonBlobSet)(nil))
var __VDLType4 *vdl.Type = vdl.TypeOf([]*vom.RawBytes(nil))
var __VDLType6 *vdl.Type = vdl.TypeOf(map[string]struct{}(nil))
var __VDLType2 *vdl.Type = vdl.TypeOf(map[nosql.BlobRef]struct{}(nil))
var __VDLType_v_io_v23_services_syncbase_nosql_BlobRef *vdl.Type = vdl.TypeOf(nosql.BlobRef(""))
var __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobAny *vdl.Type = vdl.TypeOf(BlobAny{})
var __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobInfo *vdl.Type = vdl.TypeOf(BlobInfo{})
var __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobOpt *vdl.Type = vdl.TypeOf(BlobOpt{})
var __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobSet *vdl.Type = vdl.TypeOf(BlobSet{})
var __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_BlobUnion *vdl.Type = vdl.TypeOf(BlobUnion(BlobUnionNum{int32(0)}))
var __VDLType_v_io_x_ref_services_syncbase_vsync_testdata_NonBlobSet *vdl.Type = vdl.TypeOf(NonBlobSet{})

func __VDLEnsureNativeBuilt() {
}
