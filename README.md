# Unionize

## A tool for generating unions in Go

A fork of [zyedidia/unionize](https://github.com/zyedidia/unionize).

### Install:

```
go install github.com/spywiree/unionize@latest
```

### Changes:

- Added an option to generate tagged unions
- Fixed several bugs

### Usage:

```
Usage: unionize <input> <template-name> <output> <union-name> [flags]

A tool for generating unions in Go

Arguments:
  <input>            input package name
  <template-name>    input template type name
  <output>           output file name
  <union-name>       output union type name

Flags:
  -h, --help      Show context-sensitive help.
  -S, --strict    exit on package errors
  -T, --tagged    generate tagged union

unionize: error: expected "<input> <template-name> <output> <union-name>"
```
