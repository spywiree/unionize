package generate

import (
	"bytes"
	"embed"
	"go/format"
	"math"
	"text/template"

	"github.com/spywiree-priv/unionize/unionize/parse"
)

//go:embed tmpl
var fsys embed.FS

var tmpl = template.Must(template.ParseFS(fsys, "tmpl/*.go.tmpl"))

type UnionData struct {
	PackageName string
	Imports     []parse.PkgImport

	Name    string
	BufSize int64
	BufType string

	Tagged    bool
	NoPtrRecv bool

	Fields []parse.Field
}

func (ud *UnionData) EnumName() string {
	return ud.Name + "Type"
}

func (ud *UnionData) EnumType() string {
	switch {
	case int64(len(ud.Fields)) <= math.MaxUint8:
		return "uint8"
	case int64(len(ud.Fields)) <= math.MaxUint16:
		return "uint16"
	case int64(len(ud.Fields)) <= math.MaxUint32:
		return "uint32"
	default:
		return "uint64"
	}
}

func (ud *UnionData) GenerateUnsafe() ([]byte, error) {
	return ud.generate(tmpl.Lookup("unsafe.go.tmpl"))
}

func (ud *UnionData) GenerateSafe() ([]byte, error) {
	return ud.generate(tmpl.Lookup("safe.go.tmpl"))
}

func (ud *UnionData) generate(tmpl *template.Template) ([]byte, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, ud)
	if err != nil {
		return nil, err
	}
	// return buf.Bytes(), nil
	return format.Source(buf.Bytes())
}
