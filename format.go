package validator

import "strings"

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

func (n Number) FormatSpaces() string {
	return n.FormatSeparator(' ')
}

func (n Number) FormatDashes() string {
	return n.FormatSeparator('-')
}

func (n Number) Mask() string {
	if len(n) != 10 {
		return string(n)
	}
	return "*** *** " + string(n[6:])
}
