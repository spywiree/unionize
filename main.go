package main

import (
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/spywiree/unionize/unionize/generate"
	"github.com/spywiree/unionize/unionize/parse"
)

var cli struct {
	Input        string `arg:"" help:"input package name"`
	TemplateName string `arg:"" help:"input template type name"`
	Output       string `arg:"" type:"path" help:"output file name"`
	UnionName    string `arg:"" help:"output union type name"`

	OutputPkg string `short:"P" default:"main" help:"output package name"`
	Warn      bool   `short:"W" help:"show package import errors"`
	Tagged    bool   `short:"T" help:"generate tagged union"`
	Safe      bool   `short:"S" help:"don't use unsafe"`
	NoPtrRecv bool   `short:"R" help:"don't use pointer receiver on getters"`
}

func main() {
	_ = kong.Parse(&cli,
		kong.Name("unionize"),
		kong.Description("A tool for generating unions in Go"),
		kong.UsageOnError(),
		kong.ConfigureHelp(
			kong.HelpOptions{
				// Compact: true,
				Summary: true,
			},
		),
	)

	pkg := parse.LoadPackage(cli.Warn, cli.Input)
	u := parse.FindUnion(pkg, cli.TemplateName)
	sz, align := u.Size()
	ud := generate.UnionData{
		PackageName: cli.OutputPkg,
		Imports:     u.Imports(),

		Name:    cli.UnionName,
		BufSize: u.BufSize(sz, align),
		BufType: u.BufType(sz, align),

		Tagged:    cli.Tagged,
		NoPtrRecv: cli.NoPtrRecv,

		Fields: u.Fields(),
	}

	var data []byte
	var err error
	if !cli.Safe {
		data, err = ud.GenerateUnsafe()
	} else {
		data, err = ud.GenerateSafe()
	}
	if err != nil {
		log.Fatalln("error with union generation:", err)
	}
	err = os.WriteFile(cli.Output, data, 0600)
	if err != nil {
		log.Fatalln(err)
	}
}
