package validator

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Number
		wantErr error
	}{
		// Valid numbers (allowed formats)
		{"valid raw", "1554193974", Number("1554193974"), nil},
		{"valid dashes", "495-579-0062", Number("4955790062"), nil},
		{"valid spaces", "733 095 7763", Number("7330957763"), nil},
		{"valid dots", "739.516.3621", Number("7395163621"), nil},
		{"valid slashes", "688/609/4858", Number("6886094858"), nil},
		{"valid all zeros", "0000000000", Number("0000000000"), nil},

		// Disallowed formats
		{"mismatched separators", "688 609-4858", Number(""), ErrInvalidFormat},
		{"wrong grouping", "12-3456-7890", Number(""), ErrInvalidFormat},
		{"letters as separators", "123a456a7890", Number(""), ErrInvalidFormat},
		{"too many digits", "12345678901", Number(""), ErrInvalidFormat},
		{"empty", "", Number(""), ErrInvalidFormat},

		// Invalid checksums (valid format, bad math)
		{"invalid checksum 1", "1554193975", Number(""), ErrInvalidChecksum},
		{"invalid checksum 2", "4955790061", Number(""), ErrInvalidChecksum},

		// Invalid check digit (10)
		{"invalid check digit 10", "0074625420", Number(""), ErrInvalidChecksum},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Parse(tc.input)
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("Parse(%q) error = %v, wantErr %v", tc.input, err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("Parse(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"valid", "1554193974", true},
		{"valid formatted", "495 579 0062", true},
		{"valid dots", "733.095.7763", true},
		{"mismatched", "733.095 7763", false},
		{"invalid checksum", "1554193975", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValid(tc.input); got != tc.want {
				t.Errorf("IsValid(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
