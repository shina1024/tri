# tri
[![Go](https://github.com/shina1024/tri/actions/workflows/go.yml/badge.svg)](https://github.com/shina1024/tri/actions/workflows/go.yml)

**Ternary operator-like functions in Go**

## Overview

This Go module provides some functions that serve as alternatives to the ternary operator.

Each function evaluates a condition and, based on the result, performs one of several actions.

It can return one of two values, execute an argument-less function for deferred execution, or execute a function that can potentially return an error.

## Installation

```sh
go get github.com/shina1024/tri
```

## Usage

```go
package main

import (
    "errors"

    "github.com/shina1024/tri"
)

func main() {
    // a = "true value"
    a := tri.If(true, "true value", "false value")

    // b = "false value"
    b := tri.If(false, "true value", "false value")

    // c = "true func"
    c := tri.IfExec(true,
        func() string { return "true func" },
        func() string { return "false func" })

    // d = "true func", err = nil
    d, err := tri.IfExecWithErr(true,
        func() (string, error) { return "true func", nil },
        func() (string, error) { return "false func", errors.New("an error occurred") })
    if err != nil {
        panic(err)
    }

    // e = "false func", err = error
    e, err := tri.IfExecWithErr(false,
        func() (string, error) { return "true func", nil },
        func() (string, error) { return "false func", errors.New("an error occurred") })
    if err != nil {
        panic(err)
    }
}
```
