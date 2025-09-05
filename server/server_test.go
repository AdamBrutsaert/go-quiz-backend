package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateLobbyIntegration(t *testing.T) {
	s := New()
	ts := httptest.NewServer(s.createMux())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/lobby", "application/json", strings.NewReader(""))
	if err != nil {
		t.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	code, exists := response["code"]
	if !exists {
		t.Error("Response should contain 'code' field")
	}

	if len(code) != 5 {
		t.Errorf("Expected code length 5, got %d", len(code))
	}
}
