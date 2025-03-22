package parse

import (
	"errors"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

func LoadPackage(pattern string, mode packages.LoadMode, strict bool) (*packages.Package, error) {
	cfg := packages.Config{Mode: mode}
	pkgs, err := packages.Load(&cfg, pattern)
	if err != nil {
		return nil, err
	}

	if strict || pkgs[0].GoFiles == nil {
		if err = PkgError(pkgs[0]); err != nil {
			return nil, err
		}
	}

	return pkgs[0], nil
}

// Based on [packages.PrintErrors]
func PkgError(pkg *packages.Package) error {
	var sb strings.Builder
	for _, err := range pkg.Errors {
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}

	// Print pkg.Module.Error once if present.
	if pkg.Module != nil && pkg.Module.Error != nil {
		sb.WriteString(pkg.Module.Error.Err)
		sb.WriteString("\n")
	}

	if sb.Len() == 0 {
		return nil
	}
	return errors.New(sb.String())
}

func FindAndParse(pkg *packages.Package, name string) (*Union, error) {
	for id, def := range pkg.TypesInfo.Defs {
		if def == nil || id.Name != name {
			continue
		}

		s, ok := def.Type().Underlying().(*types.Struct)
		if ok {
			return Parse(pkg, s), nil
		}
	}

	return nil, errors.New("could not find struct to unionize")
}
