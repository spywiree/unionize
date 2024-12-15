package main

import (
	"bytes"
	htmltemplate "html/template"
	texttemplate "text/template"
)

type Dummy struct{}

//go:generate unionize . template union.go Union -W -T
//go:generate unionize github.com/spywiree-priv/unionize/example template union_safe.go UnionSafe -W -T -S
//go:generate unionize template.go template union_safe.go UnionSafe -W -T -S
//nolint:all
type template struct {
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
