package validator

import "strings"

// FormatSeparator returns n grouped as NNN<s>NNN<s>NNNN using the separator s.
// If n is not 10 digits it is returned unchanged.
func (n Number) FormatSeparator(s byte) string {
	if len(n) != 10 {
		return string(n)
	}

	var result strings.Builder

	for i, char := range n {
		if i == 3 || i == 6 {
			result.WriteByte(s)
		}
		result.WriteRune(char)
	}

	return result.String()
}

// FormatSpaces returns n grouped as "NNN NNN NNNN".
func (n Number) FormatSpaces() string {
	return n.FormatSeparator(' ')
}

// FormatDashes returns n grouped as "NNN-NNN-NNNN".
func (n Number) FormatDashes() string {
	return n.FormatSeparator('-')
}

// Mask returns n with all but the final four digits replaced, as "*** *** NNNN".
// If n is not 10 digits it is returned unchanged.
func (n Number) Mask() string {
	if len(n) != 10 {
		return string(n)
	}
	return "*** *** " + string(n[6:])
}
