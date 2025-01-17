package parse

import (
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
)

type Union struct {
	pkg      *packages.Package
	template *types.Struct
}

// I believe there are no types in Go with alignment 16.
var alignments = map[int64]string{
	1: "uint8",
	2: "uint16",
	4: "uint32",
	8: "uint64",
}

// Size returns the size and alignment necessary for the underyling union
// buffer given the template struct.
func (u Union) Size() (int64, int64) {
	if u.template.NumFields() == 0 {
		log.Fatalln("error: union template is empty")
	}

	var maxsz int64
	var maxalign int64
	for i := 0; i < u.template.NumFields(); i++ {
		f := u.template.Field(i)
		sz := u.pkg.TypesSizes.Sizeof(f.Type())
		align := u.pkg.TypesSizes.Alignof(f.Type())
		if sz > maxsz {
			maxsz = sz
		}
		if align > maxalign {
			maxalign = align
		}
	}
	if maxsz%maxalign != 0 {
		maxsz = maxsz - maxsz%maxalign + maxalign
	}

	if _, ok := alignments[maxalign]; !ok {
		log.Printf("warning: alignment of %d cannot be satisfied with a primitive type, using alignment of %d instead\n", maxalign, 8)
		maxalign = 8
	}

	return maxsz, maxalign
}

func (u Union) BufSize(sz, align int64) int64 {
	return sz / align
}

func (u Union) BufType(sz, align int64) string {
	return alignments[align]
}

func (u Union) Imports() []PkgImport {
	x := NewImporter(u.pkg.Types)
	for i := range u.template.NumFields() {
		x.GetTypeImports(u.template.Field(i).Type())
	}
	return x.Imports
}

func qual(pkg *types.Package) types.Qualifier {
	if pkg == nil {
		return nil
	}
	return func(other *types.Package) string {
		if pkg == other {
			return ""
		}
		return other.Name()
	}
}

func (u Union) Fields() []Field {
	fields := make([]Field, u.template.NumFields())
	for i := range u.template.NumFields() {
		fields[i] = Field{
			Name: u.template.Field(i).Name(),
			Type: types.TypeString(
				u.template.Field(i).Type(),
				qual(u.pkg.Types),
			),
		}
	}
	return fields
}
