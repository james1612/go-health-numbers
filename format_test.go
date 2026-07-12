package validator

import (
	"testing"
)

func Test_Number_FormatSeparator(t *testing.T) {
	tests := []struct {
		name      string
		number    Number
		separator byte
		want      string
	}{
		{"slash separator", Number("4955790062"), '/', "495/579/0062"},
		{"dash separator", Number("4955790062"), '-', "495-579-0062"},
		{"space separator", Number("4955790062"), ' ', "495 579 0062"},
		{"pipe separator", Number("1554193974"), '|', "155|419|3974"},
		{"dot separator", Number("7330957763"), '.', "733.095.7763"},
		{"underscore separator", Number("0123456789"), '_', "012_345_6789"},
		{"comma separator", Number("6886094858"), ',', "688,609,4858"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.number.FormatSeparator(tc.separator)
			if got != tc.want {
				t.Errorf("FormatSeparator(%q) = %q, want %q", tc.separator, got, tc.want)
			}
		})
	}
}

func Test_Number_FormatSpaces(t *testing.T) {
	tests := []struct {
		name   string
		number Number
		want   string
	}{
		{"valid number 1", Number("4955790062"), "495 579 0062"},
		{"valid number 2", Number("1554193974"), "155 419 3974"},
		{"valid number 3", Number("7330957763"), "733 095 7763"},
		{"valid number 4", Number("0123456789"), "012 345 6789"},
		{"all zeros", Number("0000000000"), "000 000 0000"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.number.FormatSpaces()
			if got != tc.want {
				t.Errorf("FormatSpaces() = %q, want %q", got, tc.want)
			}
		})
	}
}

func Test_Number_FormatDashes(t *testing.T) {
	tests := []struct {
		name   string
		number Number
		want   string
	}{
		{"valid number 1", Number("4955790062"), "495-579-0062"},
		{"valid number 2", Number("1554193974"), "155-419-3974"},
		{"valid number 3", Number("7330957763"), "733-095-7763"},
		{"valid number 4", Number("0123456789"), "012-345-6789"},
		{"all zeros", Number("0000000000"), "000-000-0000"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.number.FormatDashes()
			if got != tc.want {
				t.Errorf("FormatDashes() = %q, want %q", got, tc.want)
			}
		})
	}
}

func Test_Number_Mask(t *testing.T) {
	tests := []struct {
		name   string
		number Number
		want   string
	}{
		{"valid number 1", Number("4955790062"), "*** *** 0062"},
		{"valid number 2", Number("1554193974"), "*** *** 3974"},
		{"too short", Number("12345"), "12345"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.number.Mask()
			if got != tc.want {
				t.Errorf("Mask() = %q, want %q", got, tc.want)
			}
		})
	}
}
