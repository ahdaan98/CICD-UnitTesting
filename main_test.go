package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enescakir/emoji"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(rr,req)

	if status := http.StatusOK; rr.Code != status {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := fmt.Sprintf("Hello, world! %v", emoji.Popcorn)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
