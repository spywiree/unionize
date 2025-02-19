package parse

import "go/types"

func TypeImports(
	typ types.Type,
	pkg *types.Package,
	imports map[string]*types.Package,
) {
	process := func(typ types.Type) {
		TypeImports(typ, pkg, imports)
	}
	insert := func(pkg *types.Package) {
		if pkg != nil {
			imports[pkg.Path()] = pkg
		}
	}

	// https://github.com/golang/example/tree/master/gotypes#types
	switch typ := typ.(type) {
	case *types.Basic:
		return
	case *types.Pointer:
		process(typ.Elem())
	case *types.Array:
		process(typ.Elem())
	case *types.Slice:
		process(typ.Elem())
	case *types.Map:
		process(typ.Key())
		process(typ.Elem())
	case *types.Chan:
		process(typ.Elem())
	case *types.Struct:
		for field := range typ.Fields() {
			process(field.Type())
		}
	case *types.Tuple:
		for v := range typ.Variables() {
			process(v.Type())
		}
	case *types.Signature:
		process(typ.Params())
		process(typ.Results())
		for tparam := range typ.TypeParams().TypeParams() {
			process(tparam)
		}
	case *types.Alias:
		insert(typ.Obj().Pkg())
	case *types.Named:
		insert(typ.Obj().Pkg())
	case *types.Interface:
		for e := range typ.EmbeddedTypes() {
			process(e)
		}
		for m := range typ.Methods() {
			process(m.Type())
		}
	case *types.Union:
		for term := range typ.Terms() {
			process(term.Type())
		}
	case *types.TypeParam:
		insert(typ.Obj().Pkg())
	}
}
