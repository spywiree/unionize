package parse

import (
	"fmt"
	"go/types"
	"iter"
	"strconv"
	"unicode"

	"github.com/vishalkuo/bimap"
	"golang.org/x/tools/go/packages"
)

type Union struct {
	Imports  map[string]*types.Package
	pkg      string
	pkgNames *bimap.BiMap[string, string]

	Members     []Member
	Size, Align int64
}

type Import struct {
	Name, Path string
}

type Member struct {
	Name string
	Type types.Type
}

func (m *Member) TypeString(u *Union) string {
	var qf types.Qualifier
	if u.pkg != "" {
		qf = func(other *types.Package) string {
			if u.pkg == other.Path() {
				return ""
			}
			name, _ := u.pkgNames.Get(other.Path())
			return name
		}
	}
	return types.TypeString(m.Type, qf)
}

func (u *Union) ImportsIter() iter.Seq[Import] {
	return func(yield func(Import) bool) {
		for path := range u.Imports {
			if u.pkg == path {
				continue
			}
			var i Import
			i.Name, _ = u.pkgNames.Get(path)
			i.Path = path
			if !yield(i) {
				return
			}
		}
	}
}

var alignments = map[int64]string{
	1: "uint8",
	2: "uint16",
	4: "uint32",
	8: "uint64",
}

func (u *Union) BufferType() string {
	return fmt.Sprintf("[%d]%s", u.Size/u.Align, alignments[u.Align])
}

func (u *Union) SetPackage(pkg *types.Package) error {
	u.pkg = pkg.Path()
	clear(u.pkgNames.GetForwardMap())
	clear(u.pkgNames.GetInverseMap())

	nameExists := func(name string) bool {
		return u.pkgNames.ExistsInverse(name) ||
			pkg.Scope().Lookup(name) != nil ||
			pkg.Name() == name
	}
	for _, pkg := range u.Imports {
		if u.pkg == pkg.Path() {
			continue
		} else if pkg.Name() == "main" {
			return fmt.Errorf(
				`import "%s" is a program, not an importable package`,
				pkg.Path(),
			)
		}

		name := pkg.Name()
		for i := uint64(1); nameExists(name); i++ {
			name = pkg.Name() + strconv.FormatUint(i, 10)
		}
		u.pkgNames.Insert(pkg.Path(), name)
	}

	return nil
}

func Parse(pkg *packages.Package, typ *types.Struct) *Union {
	var u Union
	u.pkgNames = bimap.NewBiMap[string, string]()

	u.Imports = make(map[string]*types.Package)
	for field := range typ.Fields() {
		TypeImports(field.Type(), pkg.Types, u.Imports)
	}

	u.Members = make([]Member, 0, typ.NumFields())
	for field := range typ.Fields() {
		u.Members = append(u.Members, Member{
			Name: mapFirstRune(field.Name(), unicode.ToUpper),
			Type: field.Type(),
		})
	}

	for field := range typ.Fields() {
		u.Size = max(u.Size, pkg.TypesSizes.Sizeof(field.Type()))
		u.Align = max(u.Align, pkg.TypesSizes.Alignof(field.Type()))
	}
	if u.Size%u.Align != 0 {
		u.Size = u.Size - u.Size%u.Align + u.Align
	}
	if _, ok := alignments[u.Align]; !ok {
		u.Align = 8
	}

	return &u
}
