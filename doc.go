// Package validator parses, validates, and formats UK NHS numbers.
//
// An NHS number is a 10-digit identifier whose final digit is a Modulus 11
// check digit. Use [Parse] to normalise and validate input into a [Number],
// or [IsValid] for a boolean check. A [Number] can be re-formatted with the
// Format methods or redacted with [Number.Mask].
package validator
