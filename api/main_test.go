package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIndex(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("incorrect status code, got %d, want %d\n", w.Code, http.StatusOK)
		t.FailNow()
	}
	if w.Body.Len() == 0 {
		t.Error("Body cannot be empty")
		t.FailNow()
	}

	headers := w.Result().Header
	if headers.Get("Access-Control-Allow-Methods") == "" {
		t.Errorf("Access-Control-Allow-Methods header cannot be empty, must be %q\n", "POST, OPTIONS")
		t.FailNow()
	}
	if headers.Get("Content-Type") != "text/html; charset=utf-8" {
		t.Errorf("incorrect Content-Type, got %q, want %q\n", headers.Get("Content-Type"), "text/html; charset=utf-8")
		t.FailNow()
	}
}
