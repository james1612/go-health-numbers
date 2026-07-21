package validator

// Number is a validated 10-digit NHS number.
type Number string

// String returns the string representation of the Number.
func (n Number) String() string {
	return string(n)
}
