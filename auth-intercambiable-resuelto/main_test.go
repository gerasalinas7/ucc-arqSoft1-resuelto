package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTAuthFlow(t *testing.T) {
	testAuthFlow(t, "jwt")
}

func TestAPIKeyAuthFlow(t *testing.T) {
	testAuthFlow(t, "apikey")
}

func TestInvalidPassword(t *testing.T) {
	server, err := newServer("jwt")
	if err != nil {
		t.Fatalf("newServer error: %v", err)
	}

	body := `{"username":"admin","password":"incorrecta"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func testAuthFlow(t *testing.T, mode string) {
	t.Helper()

	server, err := newServer(mode)
	if err != nil {
		t.Fatalf("newServer error: %v", err)
	}

	token := loginAndGetToken(t, server)

	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}

	if response["user"] != "admin" {
		t.Fatalf("expected user admin, got %q", response["user"])
	}

	if response["auth_type"] != mode {
		t.Fatalf("expected auth_type %q, got %q", mode, response["auth_type"])
	}
}

func loginAndGetToken(t *testing.T, server http.Handler) string {
	t.Helper()

	body := `{"username":"admin","password":"admin123"}`
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}

	if response["token"] == "" {
		t.Fatal("expected token in response")
	}

	return response["token"]
}
