package validator

import "unicode"

func isSeparator(b byte) bool {
	return b == ' ' || b == '-' || b == '.' || b == '/'
}

func normalise(s string) (string, error) {
	if len(s) == 0 {
		return "", ErrInvalidFormat
	}

	// Raw 10-digit format
	if len(s) == 10 {
		for _, r := range s {
			if !unicode.IsDigit(r) {
				return "", ErrInvalidFormat
			}
		}
		return s, nil
	}

	// 3-3-4 format: NNN[sep]NNN[sep]NNNN with a consistent separator
	if len(s) == 12 && isSeparator(s[3]) && s[3] == s[7] {
		var buf [10]byte
		j := 0
		for i := 0; i < 12; i++ {
			if i == 3 || i == 7 {
				continue
			}
			if !unicode.IsDigit(rune(s[i])) {
				return "", ErrInvalidFormat
			}
			buf[j] = s[i]
			j++
		}
		return string(buf[:]), nil
	}

	return "", ErrInvalidFormat
}
