package generate

import (
	"github.com/spywiree/unionize/unionize/parse"
)

type Config struct {
	Package string
	Name    string
	Tagged  bool
}

func (ud *UnionData) EnumName() string {
	return ud.Name + "Type"
}

func (ud *UnionData) EnumMemberName(m *parse.Member) string {
	return ud.Name + m.Name
}

func (ud *UnionData) GetterName(m *parse.Member) string {
	if m.Name == "String" || m.Name == "GoString" {
		return "Get" + m.Name
	}
	return m.Name
}

func (ud *UnionData) SetterName(m *parse.Member) string {
	return "Set" + m.Name
}

func (ud *UnionData) PointerName(m *parse.Member) string {
	return m.Name + "Ptr"
}
