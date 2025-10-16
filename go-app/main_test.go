package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	expected := "Hello from Go CI/CD Pipeline!"
	if rr.Body.String() != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, rr.Body.String())
	}
}
