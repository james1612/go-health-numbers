package validator

type Number string

// String returns the string representation of the Number.
func (n Number) String() string {
	return string(n)
}