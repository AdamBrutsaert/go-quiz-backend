package server

import (
	"testing"

	"github.com/google/uuid"
)

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

func TestGenerateClientID(t *testing.T) {
	// Test that generated client IDs have the expected format
	for i := 0; i < 10; i++ {
		id := generateClientID()

		// check it's an uuid
		if _, err := uuid.Parse(id); err != nil {
			t.Errorf("Client ID is not a valid UUID: %s", id)
		}
	}

	// Test uniqueness - generate multiple IDs and verify they're different
	ids := make(map[string]bool)
	for i := 0; i < 100; i++ {
		id := generateClientID()
		if ids[id] {
			t.Errorf("Duplicate client ID generated: %s", id)
		}
		ids[id] = true
	}
}
