[![Go Report Card](https://goreportcard.com/badge/github.com/terasum/gobf?ver=1)](https://goreportcard.com/report/github.com/terasum/gobf)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/terasum/gobf)

**gobf** BrainFuck language interpreter 

## Installation

```text
$ go install github.com/terasum/gobf
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

or:

```shell
$ gobf -v helloworld.bf

op: +	inter pt: 0	ds: [0]	out: []
op: +	inter pt: 0	ds: [1]	out: []
...
op: +	inter t: 2	ds: [0 87 100 32 10 0 0 0]	out: [72 101 108 108 111 32 87 111 114 108 100]
op: +	inter pt: 3	ds: [0 87 100 32 10 0 0 0]	out: [72 101 108 108 111 32 87 111 114 108 100]
op: .	inter pt: 3	ds: [0 87 100 33 10 0 0 0]	out: [72 101 108 108 111 32 87 111 114 108 100]
op: >	inter pt: 3	ds: [0 87 100 33 10 0 0 0]	out: [72 101 108 108 111 32 87 111 114 108 100 33]
op: .	inter pt: 4	ds: [0 87 100 33 10 0 0 0]	out: [72 101 108 108 111 32 87 111 114 108 100 33]
Hello World!
```

## Acknowledgments
* Thanks [github.com/spf13/cobra](https://github.com/spf13/cobra) for its powerful CLI generator.
