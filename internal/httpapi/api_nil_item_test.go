package httpapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"task015-bloom/internal/bloom"
)

// TestAddNilItem verifies that POST /add with missing "item" field
// returns 400 instead of panicking with a nil pointer dereference.
func TestAddNilItem(t *testing.T) {
	f, err := bloom.New(1000, 0.01)
	if err != nil {
		t.Fatalf("bloom.New error: %v", err)
	}
	api := New(f)
	handler := api.Handler()

	// Send request body without "item" field.
	body := strings.NewReader(`{}`)
	req := httptest.NewRequest(http.MethodPost, "/add", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for nil item, got %d; body: %s", rec.Code, rec.Body.String())
	}
}

// TestAddNullItem verifies that POST /add with explicit null item
// returns 400 instead of panicking.
func TestAddNullItem(t *testing.T) {
	f, err := bloom.New(1000, 0.01)
	if err != nil {
		t.Fatalf("bloom.New error: %v", err)
	}
	api := New(f)
	handler := api.Handler()

	body := strings.NewReader(`{"item":null}`)
	req := httptest.NewRequest(http.MethodPost, "/add", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400 for null item, got %d; body: %s", rec.Code, rec.Body.String())
	}
}
