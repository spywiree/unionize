package main

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/spywiree/unionize/unionize/generate"
	"github.com/spywiree/unionize/unionize/parse"
	"golang.org/x/tools/go/packages"
)

type cli_t struct {
	Input        string `arg:"" help:"input package name"`
	TemplateName string `arg:"" help:"input template type name"`
	Output       string `arg:"" type:"path" help:"output file name"`
	UnionName    string `arg:"" help:"output union type name"`

	Strict bool `short:"S" help:"exit on package errors"`
	Tagged bool `short:"T" help:"generate tagged union"`
}

func (cli *cli_t) Run() error {
	src, err := parse.LoadPackage(
		cli.Input,
		packages.LoadSyntax,
		cli.Strict,
	)
	if err != nil {
		return err
	}

	dst, err := parse.LoadPackage(
		filepath.Dir(cli.Output),
		packages.LoadTypes,
		false,
	)
	if err != nil {
		return err
	}

	u, err := parse.FindAndParse(src, cli.TemplateName)
	if err != nil {
		return err
	}
	if err = u.SetPackage(dst.Types); err != nil {
		return err
	}

	ud := generate.UnionData{
		Union: u,
		Config: generate.Config{
			Package: dst.Name,
			Name:    cli.UnionName,
			Tagged:  cli.Tagged,
		},
	}

	data, err := ud.Generate()
	if err != nil {
		return err
	}

	var f *os.File
	if cli.Output != "-" {
		f, err = os.OpenFile(
			cli.Output,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0600,
		)
		if err != nil {
			return err
		}
		defer f.Close()
	} else {
		f = os.Stdout
	}
	if _, err = f.Write(data); err != nil {
		return err
	}

	return nil
}

func main() {
	var cli cli_t
	ctx := kong.Parse(&cli,
		kong.Name("unionize"),
		kong.Description("A tool for generating unions in Go"),
		kong.UsageOnError(),
	)
	ctx.FatalIfErrorf(ctx.Run())
}
