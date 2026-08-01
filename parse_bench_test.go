package validator

import "testing"

func BenchmarkParse(b *testing.B) {
	cases := []struct {
		name  string
		input string
	}{
		{"raw", "1554193974"},
		{"spaces", "733 095 7763"},
		{"dashes", "495-579-0062"},
		{"invalid checksum", "1554193975"},
		{"invalid format", "688 609-4858"},
	}

	for _, bc := range cases {
		b.Run(bc.name, func(b *testing.B) {
			for b.Loop() {
				Parse(bc.input)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	cases := []struct {
		name  string
		input string
	}{
		{"valid", "1554193974"},
		{"invalid", "1554193975"},
	}

	for _, bc := range cases {
		b.Run(bc.name, func(b *testing.B) {
			for b.Loop() {
				IsValid(bc.input)
			}
		})
	}
}
