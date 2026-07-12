package validator

import (
	"testing"
)

func Test_checksumValid(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   bool
	}{
		{"valid checksum 1", "1554193974", true},
		{"valid checksum 2", "4955790062", true},
		{"valid checksum 3", "7330957763", true},
		{"invalid checksum 1", "1554193975", false},
		{"invalid checksum 2", "4955790061", false},
		{"invalid checksum 3", "7330957762", false},
		{"wrong length short", "123456789", false},
		{"wrong length long", "12345678901", false},
		{"empty string", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := checksumValid(tc.digits)
			if got != tc.want {
				t.Errorf("checksumValid(%q) = %v, want %v", tc.digits, got, tc.want)
			}
		})
	}
}

func Test_checksumDigit(t *testing.T) {
	tests := []struct {
		name      string
		digits    string
		checksum  int
		wantError bool
	}{
		{"valid checksum 4", "155419397", 4, false},
		{"valid checksum 2", "495579006", 2, false},
		{"valid checksum 3", "733095776", 3, false},
		{"valid checksum 1", "739516362", 1, false},
		{"valid checksum 8", "688609485", 8, false},
		{"valid checksum 0 from 11", "626921220", 0, false},
		{"valid checksum 9", "012345678", 9, false},
		{"valid checksum 0", "006547524", 0, false},
		{"invalid - check digit is 10", "007462542", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checksumDigit(tt.digits)

			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error for %s, got nil", tt.digits)
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error for %s: %v", tt.digits, err)
			}

			if got != tt.checksum {
				t.Errorf("Expected %d, got %d for %s", tt.checksum, got, tt.digits)
			}
		})
	}
}
