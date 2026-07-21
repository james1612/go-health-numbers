package validator

// Parse normalises s, validates its Modulus 11 checksum, and returns the
// resulting Number. It accepts a raw 10-digit string or a 3-3-4 grouped string
// with a single consistent separator. It returns ErrInvalidFormat if s is not a
// recognised format and ErrInvalidChecksum if the check digit is wrong.
func Parse(s string) (Number, error) {
	normalised, err := normalise(s)
	if err != nil {
		return "", err
	}

	if !checksumValid(normalised) {
		return "", ErrInvalidChecksum
	}

	return Number(normalised), nil
}

// Validate reports whether s is a valid NHS number by returning the error from
// [Parse], or nil if s is valid.
func Validate(s string) error {
	_, err := Parse(s)
	return err
}

// IsValid reports whether s is a valid NHS number.
func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
