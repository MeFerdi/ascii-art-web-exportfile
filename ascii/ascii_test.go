package ascii

import (
	"testing"
)

// TestLoadBanners checks if all banner files are loaded successfully
func TestPrintAscii(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		bannerStyle string
	}{
		{"Standard banner", "Hello", "standard.txt"},
		{"Shadow banner", "World", "shadow.txt"},
		{"Non-ASCII character", "Hello\x00", "standard.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of PrintAscii
			// ...
		})
	}
}
