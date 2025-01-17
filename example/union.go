// Code generated by unionize.

package main

import (
	bytes "bytes"
	template1 "html/template"
	template0 "text/template"
	"unsafe"
)

type UnionType uint8

const (
	UnionType_Uint64      UnionType = 0
	UnionType_Int64       UnionType = 1
	UnionType_String      UnionType = 2
	UnionType_importTest1 UnionType = 3
	UnionType_importTest2 UnionType = 4
	UnionType_importTest3 UnionType = 5
	UnionType_importTest4 UnionType = 6
	UnionType_importTest5 UnionType = 7
	UnionType_importTest6 UnionType = 8
	UnionType_importTest7 UnionType = 9
)

func (x UnionType) String() string {
	switch x {
	case UnionType_Uint64:
		return "Uint64"
	case UnionType_Int64:
		return "Int64"
	case UnionType_String:
		return "String"
	case UnionType_importTest1:
		return "importTest1"
	case UnionType_importTest2:
		return "importTest2"
	case UnionType_importTest3:
		return "importTest3"
	case UnionType_importTest4:
		return "importTest4"
	case UnionType_importTest5:
		return "importTest5"
	case UnionType_importTest6:
		return "importTest6"
	case UnionType_importTest7:
		return "importTest7"
	default:
		panic("unreachable")
	}
}

func (u *Union) Type() UnionType {
	return u.typ
}

type Union struct {
	typ  UnionType
	data [8]uint64
}

func (u *Union) Uint64() uint64 {
	return *(*uint64)(unsafe.Pointer(&u.data))
}
func (u *Union) SetUint64(v uint64) {
	u.typ = 0
	*(*uint64)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) Int64() int64 {
	return *(*int64)(unsafe.Pointer(&u.data))
}
func (u *Union) SetInt64(v int64) {
	u.typ = 1
	*(*int64)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) String() string {
	return *(*string)(unsafe.Pointer(&u.data))
}
func (u *Union) SetString(v string) {
	u.typ = 2
	*(*string)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest1() []bytes.Buffer {
	return *(*[]bytes.Buffer)(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest1(v []bytes.Buffer) {
	u.typ = 3
	*(*[]bytes.Buffer)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest2() [1]bytes.Buffer {
	return *(*[1]bytes.Buffer)(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest2(v [1]bytes.Buffer) {
	u.typ = 4
	*(*[1]bytes.Buffer)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest3() struct{ bytes.Buffer } {
	return *(*struct{ bytes.Buffer })(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest3(v struct{ bytes.Buffer }) {
	u.typ = 5
	*(*struct{ bytes.Buffer })(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest4() interface{ Buffer() bytes.Buffer } {
	return *(*interface{ Buffer() bytes.Buffer })(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest4(v interface{ Buffer() bytes.Buffer }) {
	u.typ = 6
	*(*interface{ Buffer() bytes.Buffer })(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest5() Dummy {
	return *(*Dummy)(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest5(v Dummy) {
	u.typ = 7
	*(*Dummy)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest6() template0.Template {
	return *(*template0.Template)(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest6(v template0.Template) {
	u.typ = 8
	*(*template0.Template)(unsafe.Pointer(&u.data)) = v
}

func (u *Union) importTest7() template1.Template {
	return *(*template1.Template)(unsafe.Pointer(&u.data))
}
func (u *Union) setImportTest7(v template1.Template) {
	u.typ = 9
	*(*template1.Template)(unsafe.Pointer(&u.data)) = v
}
