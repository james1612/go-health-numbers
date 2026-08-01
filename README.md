# NHS Number Validator

[![Go Reference](https://pkg.go.dev/badge/github.com/james1612/go-health-numbers.svg)](https://pkg.go.dev/github.com/james1612/go-health-numbers)

A lightweight, dependency-free Go library for parsing, validating, and formatting UK NHS numbers, using the Modulus 11 checksum.

## NHS Numbers

An NHS number is a unique 10-digit identifier assigned to every patient registered with the NHS in England, Wales, and the Isle of Man. The 10th digit is a Modulus 11 check digit, calculated from the preceding nine, which allows the number to be validated without a database lookup.

The NHS requires that numbers are always displayed to the public in a 3-3-4 format, separated by spaces or hyphens (e.g. `943 476 5919` or `943-476-5919`). This library accepts both formats on input, as well as dots and slashes, and can format a parsed number accordingly.

For more detail see the [NHS number specification](https://digital.nhs.uk/services/personal-demographics-service/nhs-number).

## Installation

```bash
go get github.com/james1612/go-health-numbers
```

The package is named `validator`:

```go
import validator "github.com/james1612/go-health-numbers"
```

## Usage

```go
package main

import (
	"fmt"

	validator "github.com/james1612/go-health-numbers"
)

func main() {
	// Parse returns a typed Number and validates format + checksum.
	n, err := validator.Parse("943 476 5919")
	if err != nil {
		fmt.Println("invalid:", err)
		return
	}

	fmt.Println(n)                 // 9434765919
	fmt.Println(n.FormatSpaces())  // 943 476 5919
	fmt.Println(n.FormatDashes())  // 943-476-5919
	fmt.Println(n.Mask())          // *** *** 5919

	// Or check validity without keeping the value.
	fmt.Println(validator.IsValid("9434765919")) // true
}
```

## Performance

Zero dependencies and minimal allocations. Formatted inputs (e.g. `943 476 5919`) incur one allocation to strip separators; raw 10-digit inputs allocate nothing.

```
BenchmarkParse/raw-8              54700958    18.37 ns/op    0 B/op    0 allocs/op
BenchmarkParse/spaces-8           36045552    32.76 ns/op   16 B/op    1 allocs/op
BenchmarkParse/invalid_format-8  427783339     2.80 ns/op    0 B/op    0 allocs/op
```