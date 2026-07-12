package validator

import "errors"

var (
	// ErrInvalidChecksum is returned when the input's check digit is incorrect.
	ErrInvalidChecksum = errors.New("invalid checksum")
	// ErrInvalidCheckDigit is returned when the calculated check digit is 10 (which is invalid for NHS numbers).
	ErrInvalidCheckDigit = errors.New("invalid check digit")
	// ErrInvalidFormat is returned when the input string does not match allowed NHS number formats.
	ErrInvalidFormat = errors.New("invalid format: must be 10 digits or 3-3-4 format with consistent separators")
	// ErrInvalidCharacters is returned when the input contains characters that are not allowed.
	ErrInvalidCharacters = errors.New("number contains invalid characters")
)
