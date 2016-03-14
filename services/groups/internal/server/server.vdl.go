// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: server

package server

import (
	"fmt"
	"v.io/v23/security/access"
	"v.io/v23/services/groups"
	"v.io/v23/vdl"
)

// groupData represents the persistent state of a group. (The group name is
// persisted as the store entry key.)
type groupData struct {
	Perms   access.Permissions
	Entries map[groups.BlessingPatternChunk]struct{}
}

func (groupData) __VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/services/groups/internal/server.groupData"`
}) {
}

func (m *groupData) FillVDLTarget(t vdl.Target, tt *vdl.Type) error {
	if __VDLType_v_io_x_ref_services_groups_internal_server_groupData == nil || __VDLType0 == nil {
		panic("Initialization order error: types generated for FillVDLTarget not initialized. Consider moving caller to an init() block.")
	}
	fieldsTarget1, err := t.StartFields(tt)
	if err != nil {
		return err
	}

	keyTarget2, fieldTarget3, err := fieldsTarget1.StartField("Perms")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		if err := m.Perms.FillVDLTarget(fieldTarget3, __VDLType_v_io_v23_security_access_Permissions); err != nil {
			return err
		}
		if err := fieldsTarget1.FinishField(keyTarget2, fieldTarget3); err != nil {
			return err
		}
	}
	keyTarget4, fieldTarget5, err := fieldsTarget1.StartField("Entries")
	if err != vdl.ErrFieldNoExist && err != nil {
		return err
	}
	if err != vdl.ErrFieldNoExist {

		setTarget6, err := fieldTarget5.StartSet(__VDLType1, len(m.Entries))
		if err != nil {
			return err
		}
		for key8 := range m.Entries {
			keyTarget7, err := setTarget6.StartKey()
			if err != nil {
				return err
			}

			if err := key8.FillVDLTarget(keyTarget7, __VDLType_v_io_v23_services_groups_BlessingPatternChunk); err != nil {
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

func (m *groupData) MakeVDLTarget() vdl.Target {
	return &groupDataTarget{Value: m}
}

type groupDataTarget struct {
	Value         *groupData
	permsTarget   access.PermissionsTarget
	entriesTarget unnamed_7365745b762e696f2f7632332f73657276696365732f67726f7570732e426c657373696e675061747465726e4368756e6b20737472696e675dTarget
	vdl.TargetBase
	vdl.FieldsTargetBase
}

func (t *groupDataTarget) StartFields(tt *vdl.Type) (vdl.FieldsTarget, error) {
	if !vdl.Compatible(tt, __VDLType_v_io_x_ref_services_groups_internal_server_groupData) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType_v_io_x_ref_services_groups_internal_server_groupData)
	}
	return t, nil
}
func (t *groupDataTarget) StartField(name string) (key, field vdl.Target, _ error) {
	switch name {
	case "Perms":
		t.permsTarget.Value = &t.Value.Perms
		target, err := &t.permsTarget, error(nil)
		return nil, target, err
	case "Entries":
		t.entriesTarget.Value = &t.Value.Entries
		target, err := &t.entriesTarget, error(nil)
		return nil, target, err
	default:
		return nil, nil, fmt.Errorf("field %s not in struct %v", name, __VDLType_v_io_x_ref_services_groups_internal_server_groupData)
	}
}
func (t *groupDataTarget) FinishField(_, _ vdl.Target) error {
	return nil
}
func (t *groupDataTarget) FinishFields(_ vdl.FieldsTarget) error {

	return nil
}

// map[groups.BlessingPatternChunk]struct{}
type unnamed_7365745b762e696f2f7632332f73657276696365732f67726f7570732e426c657373696e675061747465726e4368756e6b20737472696e675dTarget struct {
	Value     *map[groups.BlessingPatternChunk]struct{}
	currKey   groups.BlessingPatternChunk
	keyTarget groups.BlessingPatternChunkTarget
	vdl.TargetBase
	vdl.SetTargetBase
}

func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f67726f7570732e426c657373696e675061747465726e4368756e6b20737472696e675dTarget) StartSet(tt *vdl.Type, len int) (vdl.SetTarget, error) {
	if !vdl.Compatible(tt, __VDLType1) {
		return nil, fmt.Errorf("type %v incompatible with %v", tt, __VDLType1)
	}
	*t.Value = make(map[groups.BlessingPatternChunk]struct{})
	return t, nil
}
func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f67726f7570732e426c657373696e675061747465726e4368756e6b20737472696e675dTarget) StartKey() (key vdl.Target, _ error) {
	t.currKey = groups.BlessingPatternChunk("")
	t.keyTarget.Value = &t.currKey
	target, err := &t.keyTarget, error(nil)
	return target, err
}
func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f67726f7570732e426c657373696e675061747465726e4368756e6b20737472696e675dTarget) FinishKey(key vdl.Target) error {
	(*t.Value)[t.currKey] = struct{}{}
	return nil
}
func (t *unnamed_7365745b762e696f2f7632332f73657276696365732f67726f7570732e426c657373696e675061747465726e4368756e6b20737472696e675dTarget) FinishSet(list vdl.SetTarget) error {
	if len(*t.Value) == 0 {
		*t.Value = nil
	}

	return nil
}

func init() {
	vdl.Register((*groupData)(nil))
}

var __VDLType0 *vdl.Type = vdl.TypeOf((*groupData)(nil))
var __VDLType1 *vdl.Type = vdl.TypeOf(map[groups.BlessingPatternChunk]struct{}(nil))
var __VDLType_v_io_v23_security_access_Permissions *vdl.Type = vdl.TypeOf(access.Permissions(nil))
var __VDLType_v_io_v23_services_groups_BlessingPatternChunk *vdl.Type = vdl.TypeOf(groups.BlessingPatternChunk(""))
var __VDLType_v_io_x_ref_services_groups_internal_server_groupData *vdl.Type = vdl.TypeOf(groupData{})

func __VDLEnsureNativeBuilt() {
}
