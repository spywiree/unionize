package parse

import (
	"unicode"
	"unicode/utf8"
)

type Field struct {
	Name, Type string
}

func (f *Field) GetterName() string {
	return f.Name
}

func (f *Field) SetterName() string {
	r, size := utf8.DecodeRune([]byte(f.Name))
	if !unicode.IsLower(r) {
		return "Set" + f.Name
	} else {
		buf := []byte("set")
		buf = utf8.AppendRune(buf, unicode.ToUpper(r))
		buf = append(buf, f.Name[size:]...)
		return string(buf)
	}
}
