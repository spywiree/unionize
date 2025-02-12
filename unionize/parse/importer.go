package parse

import (
	"go/types"
	"strconv"

	"github.com/hashicorp/go-set/v3"
)

type Importer struct {
	Imports      []PkgImport
	pkg          *types.Package
	names, paths *set.Set[string]
}

type PkgImport struct {
	Name, Path string
}

func NewImporter(pkg *types.Package) Importer {
	return Importer{
		pkg:   pkg,
		names: set.New[string](0),
		paths: set.New[string](0),
	}
}

func (x *Importer) GetTypeImports(typ types.Type) {
	// var todo types.Type
	// todo = types.NewAlias()
	// todo = types.NewArray()
	// todo = types.NewChan()
	// todo = types.NewInterfaceType()
	// todo = types.NewMap()
	// todo = types.NewNamed()
	// todo = types.NewPointer()
	// todo = types.NewSignatureType()
	// todo = types.NewSlice()
	// todo = types.NewStruct()
	// todo = types.NewTuple()
	// todo = types.NewTypeParam()
	// todo = types.NewUnion()

	switch typ := typ.(type) {
	case *types.Alias:
		x.addImport(typ.Obj().Pkg())
	case *types.Array:
		x.GetTypeImports(typ.Elem())
	case *types.Chan:
		x.GetTypeImports(typ.Elem())
	case *types.Interface:
		for i := range typ.NumEmbeddeds() {
			x.GetTypeImports(typ.EmbeddedType(i))
		}
		for i := range typ.NumMethods() {
			x.GetTypeImports(typ.Method(i).Type())
		}
	case *types.Map:
		x.GetTypeImports(typ.Key())
		x.GetTypeImports(typ.Elem())
	case *types.Named:
		x.addImport(typ.Obj().Pkg())
	case *types.Pointer:
		x.GetTypeImports(typ.Elem())
	case *types.Signature /* Func */ :
		x.GetTypeImports(typ.Params())
		x.GetTypeImports(typ.Results())
		for i := range typ.TypeParams().Len() {
			x.GetTypeImports(typ.TypeParams().At(i))
		}
	case *types.Slice:
		x.GetTypeImports(typ.Elem())
	case *types.Struct:
		for i := range typ.NumFields() {
			x.GetTypeImports(typ.Field(i).Type())
		}
	case *types.Tuple:
		for i := range typ.Len() {
			x.GetTypeImports(typ.At(i).Type())
		}
	case *types.TypeParam:
		x.addImport(typ.Obj().Pkg())
	case *types.Union:
		for i := range typ.Len() {
			x.GetTypeImports(typ.Term(i).Type())
		}
	}
}

func (x *Importer) addImport(pkg *types.Package) {
	if x.pkg == pkg || pkg == nil || x.paths.Contains(pkg.Path()) {
		return
	}

	nameExists := func(name string) bool {
		return x.names.Contains(name) || x.pkg.Scope().Lookup(name) != nil
	}
	if nameExists(pkg.Name()) {
		for i := uint64(0); true; i++ {
			name := pkg.Name() + strconv.FormatUint(i, 10)
			if !nameExists(name) {
				pkg.SetName(name)
				break
			}
		}
	}

	x.names.Insert(pkg.Name())
	x.paths.Insert(pkg.Path())
	x.Imports = append(x.Imports, PkgImport{Name: pkg.Name(), Path: pkg.Path()})
}
