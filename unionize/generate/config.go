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
	switch m.Name {
	case "String", "GoString", "Type", "Interface":
		return "Get" + m.Name
	default:
		return m.Name
	}
}

func (ud *UnionData) SetterName(m *parse.Member) string {
	return "Set" + m.Name
}

func (ud *UnionData) PointerName(m *parse.Member) string {
	return m.Name + "Ptr"
}

func (ud *UnionData) OkName(m *parse.Member) string {
	return m.Name + "Ok"
}
