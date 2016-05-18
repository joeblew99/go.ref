// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: blobtestsvdl

package blobtestsvdl

import (
	"fmt"
	"v.io/v23/services/syncbase"
	"v.io/v23/vdl"
	"v.io/v23/vom"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type BlobInfo struct {
	Info string
	Br   syncbase.BlobRef
}

func (BlobInfo) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobInfo"`
}) {
}

func (x BlobInfo) VDLIsZero() bool {
	return x == BlobInfo{}
}

func (x BlobInfo) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.Br != "" {
		if err := enc.NextField("Br"); err != nil {
			return err
		}
		if err := x.Br.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobInfo) VDLRead(dec vdl.Decoder) error {
	*x = BlobInfo{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Info":
			if err := dec.StartValue(vdl.StringType); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Br":
			if err := x.Br.VDLRead(dec); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
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
		VDLIsZero() bool
		VDLWrite(vdl.Encoder) error
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

func (x BlobUnionBi) Index() int                      { return 1 }
func (x BlobUnionBi) Interface() interface{}          { return x.Value }
func (x BlobUnionBi) Name() string                    { return "Bi" }
func (x BlobUnionBi) __VDLReflect(__BlobUnionReflect) {}

func (x BlobUnionNum) VDLIsZero() bool {
	return x.Value == 0
}

func (x BlobUnionBi) VDLIsZero() bool {
	return false
}

func (x BlobUnionNum) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_union_3); err != nil {
		return err
	}
	if err := enc.NextField("Num"); err != nil {
		return err
	}
	if err := enc.StartValue(vdl.Int32Type); err != nil {
		return err
	}
	if err := enc.EncodeInt(int64(x.Value)); err != nil {
		return err
	}
	if err := enc.FinishValue(); err != nil {
		return err
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x BlobUnionBi) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_union_3); err != nil {
		return err
	}
	if err := enc.NextField("Bi"); err != nil {
		return err
	}
	if err := x.Value.VDLWrite(enc); err != nil {
		return err
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func VDLReadBlobUnion(dec vdl.Decoder, x *BlobUnion) error {
	if err := dec.StartValue(__VDLType_union_3); err != nil {
		return err
	}
	f, err := dec.NextField()
	if err != nil {
		return err
	}
	switch f {
	case "Num":
		var field BlobUnionNum
		if err := dec.StartValue(vdl.Int32Type); err != nil {
			return err
		}
		tmp, err := dec.DecodeInt(32)
		if err != nil {
			return err
		}
		field.Value = int32(tmp)
		if err := dec.FinishValue(); err != nil {
			return err
		}
		*x = field
	case "Bi":
		var field BlobUnionBi
		if err := field.Value.VDLRead(dec); err != nil {
			return err
		}
		*x = field
	case "":
		return fmt.Errorf("missing field in union %T, from %v", x, dec.Type())
	default:
		return fmt.Errorf("field %q not in union %T, from %v", f, x, dec.Type())
	}
	switch f, err := dec.NextField(); {
	case err != nil:
		return err
	case f != "":
		return fmt.Errorf("extra field %q in union %T, from %v", f, x, dec.Type())
	}
	return dec.FinishValue()
}

type BlobSet struct {
	Info string
	Bs   map[syncbase.BlobRef]struct{}
}

func (BlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobSet"`
}) {
}

func (x BlobSet) VDLIsZero() bool {
	if x.Info != "" {
		return false
	}
	if len(x.Bs) != 0 {
		return false
	}
	return true
}

func (x BlobSet) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_4); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.Bs) != 0 {
		if err := enc.NextField("Bs"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_set_1(enc, x.Bs); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_set_1(enc vdl.Encoder, x map[syncbase.BlobRef]struct{}) error {
	if err := enc.StartValue(__VDLType_set_5); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := key.VDLWrite(enc); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobSet) VDLRead(dec vdl.Decoder) error {
	*x = BlobSet{}
	if err := dec.StartValue(__VDLType_struct_4); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Info":
			if err := dec.StartValue(vdl.StringType); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Bs":
			if err := __VDLReadAnon_set_1(dec, &x.Bs); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_set_1(dec vdl.Decoder, x *map[syncbase.BlobRef]struct{}) error {
	if err := dec.StartValue(__VDLType_set_5); err != nil {
		return err
	}
	var tmpMap map[syncbase.BlobRef]struct{}
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[syncbase.BlobRef]struct{}, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		}
		var key syncbase.BlobRef
		{
			if err := key.VDLRead(dec); err != nil {
				return err
			}
		}
		if tmpMap == nil {
			tmpMap = make(map[syncbase.BlobRef]struct{})
		}
		tmpMap[key] = struct{}{}
	}
}

type BlobAny struct {
	Info string
	Baa  []*vom.RawBytes
}

func (BlobAny) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobAny"`
}) {
}

func (x BlobAny) VDLIsZero() bool {
	if x.Info != "" {
		return false
	}
	if len(x.Baa) != 0 {
		return false
	}
	return true
}

func (x BlobAny) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.Baa) != 0 {
		if err := enc.NextField("Baa"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, x.Baa); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_2(enc vdl.Encoder, x []*vom.RawBytes) error {
	if err := enc.StartValue(__VDLType_list_7); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for i := 0; i < len(x); i++ {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if x[i] == nil {
			if err := enc.NilValue(vdl.AnyType); err != nil {
				return err
			}
		} else {
			if err := x[i].VDLWrite(enc); err != nil {
				return err
			}
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobAny) VDLRead(dec vdl.Decoder) error {
	*x = BlobAny{}
	if err := dec.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Info":
			if err := dec.StartValue(vdl.StringType); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Baa":
			if err := __VDLReadAnon_list_2(dec, &x.Baa); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_2(dec vdl.Decoder, x *[]*vom.RawBytes) error {
	if err := dec.StartValue(__VDLType_list_7); err != nil {
		return err
	}
	switch len := dec.LenHint(); {
	case len > 0:
		*x = make([]*vom.RawBytes, 0, len)
	default:
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		}
		var elem *vom.RawBytes
		elem = new(vom.RawBytes)
		if err := elem.VDLRead(dec); err != nil {
			return err
		}
		*x = append(*x, elem)
	}
}

type NonBlobSet struct {
	Info string
	S    map[string]struct{}
}

func (NonBlobSet) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.NonBlobSet"`
}) {
}

func (x NonBlobSet) VDLIsZero() bool {
	if x.Info != "" {
		return false
	}
	if len(x.S) != 0 {
		return false
	}
	return true
}

func (x NonBlobSet) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_8); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if len(x.S) != 0 {
		if err := enc.NextField("S"); err != nil {
			return err
		}
		if err := __VDLWriteAnon_set_3(enc, x.S); err != nil {
			return err
		}
	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_set_3(enc vdl.Encoder, x map[string]struct{}) error {
	if err := enc.StartValue(__VDLType_set_9); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(key); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *NonBlobSet) VDLRead(dec vdl.Decoder) error {
	*x = NonBlobSet{}
	if err := dec.StartValue(__VDLType_struct_8); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Info":
			if err := dec.StartValue(vdl.StringType); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "S":
			if err := __VDLReadAnon_set_3(dec, &x.S); err != nil {
				return err
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_set_3(dec vdl.Decoder, x *map[string]struct{}) error {
	if err := dec.StartValue(__VDLType_set_9); err != nil {
		return err
	}
	var tmpMap map[string]struct{}
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string]struct{}, len)
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		}
		var key string
		{
			if err := dec.StartValue(vdl.StringType); err != nil {
				return err
			}
			var err error
			if key, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		}
		if tmpMap == nil {
			tmpMap = make(map[string]struct{})
		}
		tmpMap[key] = struct{}{}
	}
}

type BlobOpt struct {
	Info string
	Bo   *BlobInfo
}

func (BlobOpt) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/syncbase/vsync/testdata.BlobOpt"`
}) {
}

func (x BlobOpt) VDLIsZero() bool {
	return x == BlobOpt{}
}

func (x BlobOpt) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_10); err != nil {
		return err
	}
	if x.Info != "" {
		if err := enc.NextField("Info"); err != nil {
			return err
		}
		if err := enc.StartValue(vdl.StringType); err != nil {
			return err
		}
		if err := enc.EncodeString(x.Info); err != nil {
			return err
		}
		if err := enc.FinishValue(); err != nil {
			return err
		}
	}
	if x.Bo != nil {
		if err := enc.NextField("Bo"); err != nil {
			return err
		}
		enc.SetNextStartValueIsOptional()

		if err := x.Bo.VDLWrite(enc); err != nil {
			return err
		}

	}
	if err := enc.NextField(""); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *BlobOpt) VDLRead(dec vdl.Decoder) error {
	*x = BlobOpt{}
	if err := dec.StartValue(__VDLType_struct_10); err != nil {
		return err
	}
	for {
		f, err := dec.NextField()
		if err != nil {
			return err
		}
		switch f {
		case "":
			return dec.FinishValue()
		case "Info":
			if err := dec.StartValue(vdl.StringType); err != nil {
				return err
			}
			var err error
			if x.Info, err = dec.DecodeString(); err != nil {
				return err
			}
			if err := dec.FinishValue(); err != nil {
				return err
			}
		case "Bo":
			if err := dec.StartValue(__VDLType_optional_11); err != nil {
				return err
			}
			if dec.IsNil() {
				x.Bo = nil
				if err := dec.FinishValue(); err != nil {
					return err
				}
			} else {
				x.Bo = new(BlobInfo)
				dec.IgnoreNextStartValue()
				if err := x.Bo.VDLRead(dec); err != nil {
					return err
				}
			}
		default:
			if err := dec.SkipValue(); err != nil {
				return err
			}
		}
	}
}

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_struct_1    *vdl.Type
	__VDLType_string_2    *vdl.Type
	__VDLType_union_3     *vdl.Type
	__VDLType_struct_4    *vdl.Type
	__VDLType_set_5       *vdl.Type
	__VDLType_struct_6    *vdl.Type
	__VDLType_list_7      *vdl.Type
	__VDLType_struct_8    *vdl.Type
	__VDLType_set_9       *vdl.Type
	__VDLType_struct_10   *vdl.Type
	__VDLType_optional_11 *vdl.Type
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

	// Register types.
	vdl.Register((*BlobInfo)(nil))
	vdl.Register((*BlobUnion)(nil))
	vdl.Register((*BlobSet)(nil))
	vdl.Register((*BlobAny)(nil))
	vdl.Register((*NonBlobSet)(nil))
	vdl.Register((*BlobOpt)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*BlobInfo)(nil)).Elem()
	__VDLType_string_2 = vdl.TypeOf((*syncbase.BlobRef)(nil))
	__VDLType_union_3 = vdl.TypeOf((*BlobUnion)(nil))
	__VDLType_struct_4 = vdl.TypeOf((*BlobSet)(nil)).Elem()
	__VDLType_set_5 = vdl.TypeOf((*map[syncbase.BlobRef]struct{})(nil))
	__VDLType_struct_6 = vdl.TypeOf((*BlobAny)(nil)).Elem()
	__VDLType_list_7 = vdl.TypeOf((*[]*vom.RawBytes)(nil))
	__VDLType_struct_8 = vdl.TypeOf((*NonBlobSet)(nil)).Elem()
	__VDLType_set_9 = vdl.TypeOf((*map[string]struct{})(nil))
	__VDLType_struct_10 = vdl.TypeOf((*BlobOpt)(nil)).Elem()
	__VDLType_optional_11 = vdl.TypeOf((*BlobInfo)(nil))

	return struct{}{}
}
