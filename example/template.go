package main

import (
	"bytes"
	htmltemplate "html/template"
	texttemplate "text/template"
)

type Dummy struct{}

//go:generate go tool github.com/spywiree/unionize . template union.go Union -W -T
//go:generate go tool github.com/spywiree/unionize . template union_safe.go UnionSafe -W -T -S
type template struct { //nolint:all
	Uint64 uint64
	Int64  int64
	String string

	importTest1 []bytes.Buffer
	importTest2 [1]bytes.Buffer
	importTest3 struct{ bytes.Buffer }
	importTest4 interface{ Buffer() bytes.Buffer }
	importTest5 Dummy
	importTest6 texttemplate.Template
	importTest7 htmltemplate.Template
}
