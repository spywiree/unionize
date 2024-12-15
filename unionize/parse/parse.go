package parse

import (
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
)

func LoadPackage(warn bool, patterns ...string) *packages.Package {
	cfg := packages.Config{Mode: packages.LoadTypes | packages.LoadSyntax | packages.LoadImports}
	pkgs, err := packages.Load(&cfg, patterns...)
	if err != nil {
		log.Fatalln("load:", err)
	}
	if warn {
		packages.PrintErrors(pkgs)
	}
	if len(pkgs) == 0 {
		log.Fatalln("error: no package found")
	}
	return pkgs[0]
}

// FindUnion finds the struct that should be used as a template for
// the union.
func FindUnion(pkg *packages.Package, name string) Union {
	for _, d := range pkg.TypesInfo.Defs {
		if d != nil && d.Name() == name {
			s, ok := d.Type().Underlying().(*types.Struct)
			if ok {
				return Union{pkg: pkg, template: s}
			}
		}
	}
	log.Fatalln("error: could not find struct to unionize")
	return Union{}
}

// // GetImports returns the names of any packages that are needed to access
// // the types in the union fields.
// func GetImports(fields []Field, pkg *types.Package) []PkgImport {
// 	i := NewImporter(pkg)

// 	for _, f := range fields {
// 		i.GetTypeImports(f.typ)
// 	}
// }
