package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HelloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := map[string]string{"message": "Hello, World!"}
	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}
	if response["message"] != expected["message"] {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response["message"], expected["message"])
	}
}
