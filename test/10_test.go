package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedBody := "Hello, World!"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, w.Body.String())
	}
}
