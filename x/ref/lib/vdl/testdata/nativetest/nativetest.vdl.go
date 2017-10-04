// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: nativetest

// Package nativetest tests a package with native type conversions.
package nativetest

import (
	"time"
	"v.io/v23/vdl"
	"v.io/v23/vdl/testdata/nativetest"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

type WireString int32

func (WireString) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireString"`
}) {
}

func (x WireString) VDLIsZero() bool {
	return x == 0
}

func (x WireString) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueInt(__VDLType_int32_1, int64(x)); err != nil {
		return err
	}
	return nil
}

func (x *WireString) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueInt(32); {
	case err != nil:
		return err
	default:
		*x = WireString(value)
	}
	return nil
}

type WireTime int32

func (WireTime) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireTime"`
}) {
}

func (x WireTime) VDLIsZero() bool {
	return x == 0
}

func (x WireTime) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueInt(__VDLType_int32_2, int64(x)); err != nil {
		return err
	}
	return nil
}

func (x *WireTime) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueInt(32); {
	case err != nil:
		return err
	default:
		*x = WireTime(value)
	}
	return nil
}

type WireSamePkg int32

func (WireSamePkg) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireSamePkg"`
}) {
}

func (x WireSamePkg) VDLIsZero() bool {
	return x == 0
}

func (x WireSamePkg) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueInt(__VDLType_int32_3, int64(x)); err != nil {
		return err
	}
	return nil
}

func (x *WireSamePkg) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueInt(32); {
	case err != nil:
		return err
	default:
		*x = WireSamePkg(value)
	}
	return nil
}

type WireMultiImport int32

func (WireMultiImport) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireMultiImport"`
}) {
}

func (x WireMultiImport) VDLIsZero() bool {
	return x == 0
}

func (x WireMultiImport) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueInt(__VDLType_int32_4, int64(x)); err != nil {
		return err
	}
	return nil
}

func (x *WireMultiImport) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueInt(32); {
	case err != nil:
		return err
	default:
		*x = WireMultiImport(value)
	}
	return nil
}

type WireRenameMe int32

func (WireRenameMe) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireRenameMe"`
}) {
}

func (x WireRenameMe) VDLIsZero() bool {
	return x == 0
}

func (x WireRenameMe) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueInt(__VDLType_int32_5, int64(x)); err != nil {
		return err
	}
	return nil
}

func (x *WireRenameMe) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueInt(32); {
	case err != nil:
		return err
	default:
		*x = WireRenameMe(value)
	}
	return nil
}

type WireAll struct {
	A string
	B time.Time
	C nativetest.NativeSamePkg
	D map[nativetest.NativeSamePkg]time.Time
	E WireRenameMe
}

func (WireAll) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.WireAll"`
}) {
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
	if err := enc.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	if x.A != "" {
		var wire WireString
		if err := WireStringFromNative(&wire, x.A); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(0, __VDLType_int32_1, int64(wire)); err != nil {
			return err
		}
	}
	if !x.B.IsZero() {
		var wire WireTime
		if err := WireTimeFromNative(&wire, x.B); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(1, __VDLType_int32_2, int64(wire)); err != nil {
			return err
		}
	}
	if x.C != "" {
		var wire WireSamePkg
		if err := WireSamePkgFromNative(&wire, x.C); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(2, __VDLType_int32_3, int64(wire)); err != nil {
			return err
		}
	}
	if x.D != nil {
		var wire WireMultiImport
		if err := WireMultiImportFromNative(&wire, x.D); err != nil {
			return err
		}
		if err := enc.NextFieldValueInt(3, __VDLType_int32_4, int64(wire)); err != nil {
			return err
		}
	}
	if x.E != 0 {
		if err := enc.NextFieldValueInt(4, __VDLType_int32_5, int64(x.E)); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireAll) VDLRead(dec vdl.Decoder) error {
	*x = WireAll{}
	if err := dec.StartValue(__VDLType_struct_6); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_6 {
			index = __VDLType_struct_6.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			var wire WireString
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireStringToNative(wire, &x.A); err != nil {
				return err
			}
		case 1:
			var wire WireTime
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireTimeToNative(wire, &x.B); err != nil {
				return err
			}
		case 2:
			var wire WireSamePkg
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireSamePkgToNative(wire, &x.C); err != nil {
				return err
			}
		case 3:
			var wire WireMultiImport
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := WireMultiImportToNative(wire, &x.D); err != nil {
				return err
			}
		case 4:
			switch value, err := dec.ReadValueInt(32); {
			case err != nil:
				return err
			default:
				x.E = WireRenameMe(value)
			}
		}
	}
}

type ignoreme string

func (ignoreme) VDLReflect(struct {
	Name string `vdl:"v.io/x/ref/lib/vdl/testdata/nativetest.ignoreme"`
}) {
}

func (x ignoreme) VDLIsZero() bool {
	return x == ""
}

func (x ignoreme) VDLWrite(enc vdl.Encoder) error {
	if err := enc.WriteValueString(__VDLType_string_7, string(x)); err != nil {
		return err
	}
	return nil
}

func (x *ignoreme) VDLRead(dec vdl.Decoder) error {
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		*x = ignoreme(value)
	}
	return nil
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

// Hold type definitions in package-level variables, for better performance.
var (
	__VDLType_int32_1  *vdl.Type
	__VDLType_int32_2  *vdl.Type
	__VDLType_int32_3  *vdl.Type
	__VDLType_int32_4  *vdl.Type
	__VDLType_int32_5  *vdl.Type
	__VDLType_struct_6 *vdl.Type
	__VDLType_string_7 *vdl.Type
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

	// Initialize type definitions.
	__VDLType_int32_1 = vdl.TypeOf((*WireString)(nil))
	__VDLType_int32_2 = vdl.TypeOf((*WireTime)(nil))
	__VDLType_int32_3 = vdl.TypeOf((*WireSamePkg)(nil))
	__VDLType_int32_4 = vdl.TypeOf((*WireMultiImport)(nil))
	__VDLType_int32_5 = vdl.TypeOf((*WireRenameMe)(nil))
	__VDLType_struct_6 = vdl.TypeOf((*WireAll)(nil)).Elem()
	__VDLType_string_7 = vdl.TypeOf((*ignoreme)(nil))

	return struct{}{}
}