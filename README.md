# termcopy

Tiny library for copying text to the system clipboard using the ANSI OSC52 sequence from within a terminal.

[![Build status](https://img.shields.io/github/workflow/status/purpleclay/termcopy/ci?style=flat-square&logo=go)](https://github.com/purpleclay/termcopy/actions?workflow=ci)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/purpleclay/termcopy?style=flat-square)](https://goreportcard.com/report/github.com/purpleclay/termcopy)
[![Go Version](https://img.shields.io/github/go-mod/go-version/purpleclay/termcopy.svg?style=flat-square)](go.mod)
[![codecov](https://codecov.io/gh/purpleclay/termcopy/branch/main/graph/badge.svg)](https://codecov.io/gh/purpleclay/termcopy)

## Quick Start

Install:

```sh
go get github.com/purpleclay/termcopy
```

Use:

```go
package main

import (
    "github.com/purpleclay/termcopy"
)

func main() {
    // Copy a stream straight to the terminal clipboard
    termcopy.Stream(os.Stdin)

    // Copy a string to the terminal clipboard
    termcopy.String("Hello there!!")

    // Copy a series of bytes to the terminal clipboard
    termcopy.Bytes([]byte("Hello there!!"))
}
```
