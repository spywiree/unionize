# Unionize

## A tool for generating unions in Go

A fork of [zyedidia/unionize](https://github.com/zyedidia/unionize).

### Install:

```
go install github.com/spywiree/unionize@latest
```

### Changes:

- Added an option to generate tagged unions.
- Added an option to disable using pointer receivers in getters.
- Fixed several bugs.

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
  -h, --help                 Show context-sensitive help.
  -P, --output-pkg="main"    output package name
  -W, --warn                 show package import errors
  -T, --tagged               generate tagged union
  -S, --safe                 don't use unsafe
  -R, --no-ptr-recv          don't use pointer receiver on getters

unionize: error: expected "<input> <template-name> <output> <union-name>"
```
