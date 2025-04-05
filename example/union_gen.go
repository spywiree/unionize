package main

import (
	"unsafe"
)

type UnionType uint8

const (
	UnionUint   UnionType = 0
	UnionInt    UnionType = 1
	UnionString UnionType = 2
)

func (x UnionType) String() string {
	switch x {
	case UnionUint:
		return "Uint"
	case UnionInt:
		return "Int"
	case UnionString:
		return "String"
	default:
		panic("unreachable")
	}
}

func (u *Union) Type() UnionType {
	return u.typ
}

type Union struct {
	typ  UnionType
	data [2]uint64
}

func (u *Union) Uint() uint64 {
	return *(*uint64)(unsafe.Pointer(&u.data))
}
func (u *Union) SetUint(v uint64) {
	u.typ = UnionUint
	*(*uint64)(unsafe.Pointer(&u.data)) = v
}
func (u *Union) UintPtr() *uint64 {
	return (*uint64)(unsafe.Pointer(&u.data))
}
func (u *Union) UintOk() (uint64, bool) {
	if u.Type() == UnionUint {
		return u.Uint(), true
	}
	var zero uint64
	return zero, false
}

func (u *Union) Int() int64 {
	return *(*int64)(unsafe.Pointer(&u.data))
}
func (u *Union) SetInt(v int64) {
	u.typ = UnionInt
	*(*int64)(unsafe.Pointer(&u.data)) = v
}
func (u *Union) IntPtr() *int64 {
	return (*int64)(unsafe.Pointer(&u.data))
}
func (u *Union) IntOk() (int64, bool) {
	if u.Type() == UnionInt {
		return u.Int(), true
	}
	var zero int64
	return zero, false
}

func (u *Union) GetString() string {
	return *(*string)(unsafe.Pointer(&u.data))
}
func (u *Union) SetString(v string) {
	u.typ = UnionString
	*(*string)(unsafe.Pointer(&u.data)) = v
}
func (u *Union) StringPtr() *string {
	return (*string)(unsafe.Pointer(&u.data))
}
func (u *Union) StringOk() (string, bool) {
	if u.Type() == UnionString {
		return u.GetString(), true
	}
	var zero string
	return zero, false
}

func (u *Union) Interface() any {
	switch u.Type() {
	case UnionUint:
		return u.Uint()
	case UnionInt:
		return u.Int()
	case UnionString:
		return u.GetString()
	default:
		panic("unreachable")
	}
}
