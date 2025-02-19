package generate

import (
	"bytes"
	_ "embed"
	"go/format"
	"math"
	"text/template"

	"github.com/spywiree/unionize/unionize/parse"
)

type UnionData struct {
	*parse.Union
	Config
}

func (ud *UnionData) EnumType() string {
	switch {
	case int64(len(ud.Members)) <= math.MaxUint8:
		return "uint8"
	case int64(len(ud.Members)) <= math.MaxUint16:
		return "uint16"
	case int64(len(ud.Members)) <= math.MaxUint32:
		return "uint32"
	default:
		return "uint64"
	}
}

//go:embed union.go.tmpl
var unionTmpl string

var union = template.Must(template.New("union").Parse(unionTmpl))

func (ud *UnionData) Generate() ([]byte, error) {
	var buf bytes.Buffer
	err := union.Execute(&buf, ud)
	if err != nil {
		return nil, err
	}
	return format.Source(buf.Bytes())
}
