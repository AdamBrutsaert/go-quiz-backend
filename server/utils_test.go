package server

import "testing"

func TestGenerateLobbyCode(t *testing.T) {
	s := New()

	// Test that generated codes are 5 characters long
	for i := 0; i < 100; i++ {
		code := s.generateLobbyCode()
		if len(code) != 5 {
			t.Errorf("Expected code length 5, got %d for code '%s'", len(code), code)
		}

		// Verify code contains only alphanumeric characters
		for _, char := range code {
			if !((char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
				t.Errorf("Code contains invalid character: %c in code '%s'", char, code)
			}
		}
	}
}
