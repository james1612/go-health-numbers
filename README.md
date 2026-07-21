# NHS Number Validator

A lightweight, dependency-free Go library for parsing, validating, and formatting UK NHS numbers, using the Modulus 11 checksum.

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

## Validation Rules

An input is accepted when it meets all the following:

- It is either a raw 10-digit string (`NNNNNNNNNN`) or a 3-3-4 grouped string
  using a single, consistent separator — space, dash, dot, or slash
  (e.g. `943 476 5919`, `943-476-5919`).
- After removing separators it contains exactly 10 numeric digits.
- It passes the official NHS **Modulus 11** checksum:
    1. Multiply each of the first 9 digits by its position weight (10, 9, 8, ..., 2).
    2. Sum these products.
    3. Compute the check digit as `11 - (sum % 11)`; a result of 11 becomes 0.
    4. A result of 10 is invalid (such a number is never issued).
    5. Compare the check digit with the 10th digit.