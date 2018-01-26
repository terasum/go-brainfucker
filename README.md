[![Go Report Card](https://goreportcard.com/badge/github.com/terasum/gobf)](https://goreportcard.com/report/github.com/terasum/gobf)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/terasum/gobf)

**gobf** BrainFuck language interpreter 

## Installation

```text
$ go get github.com/terasum/gobf
$ gobf --help
gobf is a command line interpreter for BrainFuck language.
you can use:
    gobf source.bf
to interpreter the BrainFuck code

Usage:
  gobf [flags]

Flags:
  -v, --verbose   verbose output

```

## Example

helloworld.bf
```
++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.
```

You can run this code by:
```shell
$ gobf helloworld.bf
Hello World!
```

## Acknowledgments
* Thanks [github.com/spf13/cobra](https://github.com/spf13/cobra) for its powerful CLI generator.
