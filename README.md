# go-foreground

go-foreground is a Go package that provides the PID and title of the foreground window on Windows, Linux (X11) and macOS.

[![Reference](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/LightningDev1/go-foreground)
[![Linter](https://goreportcard.com/badge/github.com/LightningDev1/go-foreground?style=flat-square)](https://goreportcard.com/report/github.com/LightningDev1/go-foreground)
[![Build status](https://github.com/LightningDev1/go-foreground/actions/workflows/ci.yml/badge.svg)](https://github.com/LightningDev1/go-foreground/actions)

## Installation

```sh
go get github.com/LightningDev1/go-foreground
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/LightningDev1/go-foreground"
)

func main() {
    pid, _ := foreground.GetForegroundPID()
    title, _ := foreground.GetForegroundTitle()

    fmt.Printf("PID: %d\ntitle: %s\n", pid, title)
}
```
