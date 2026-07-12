package validator

import (
	"errors"
)

// checksumValid reports whether the provided string is a valid NHS number using the Modulus 11 algorithm.
// It assumes the input is already validated as a 10-digit numeric string.
func checksumValid(digits string) bool {
	if len(digits) != 10 {
		return false
	}

	checksum, err := checksumDigit(digits[0:9])
	if err != nil {
		return false
	}

	lastDigit := int(digits[9] - '0')

	return checksum == lastDigit
}

// checksumDigit calculates the check digit for the first 9 digits of an NHS number.
func checksumDigit(first9 string) (int, error) {
	if len(first9) != 9 {
		return 0, errors.New("must be 9 digits")
	}

	total := 0
	for i := 0; i < 9; i++ {
		digit := int(first9[i] - '0')
		multiplier := 10 - i
		total += digit * multiplier
	}

	remainder := total % 11
	checkDigit := 11 - remainder

	if checkDigit == 11 {
		return 0, nil
	}

	if checkDigit == 10 {
		return 0, errors.New("invalid number")
	}

	return checkDigit, nil
}
