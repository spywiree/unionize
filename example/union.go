package main

//go:generate go tool github.com/spywiree/unionize . union union_gen.go Union -T
type union struct { //nolint:all
	uint uint64
	Int  int64
	string
}
