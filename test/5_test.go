package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	HelloHandler(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Unexpected status code: %d", status)
	}

	body := rr.Body.String()
	if body != "Hello, World!" {
		t.Errorf("Unexpected response body: %s", body)
	}
}
