package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestHandleSuccess(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler := NewHealthCheckHandler()

	handler.Handle(w, r)

	results := w.Result()

	expect := 200
	got := results.StatusCode

	if expect != got {
		t.Errorf("Expect %d received %d", expect, got)
	}
}

func TestHandleNotAllowed(t *testing.T) {
	r := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()

	handler := NewHealthCheckHandler()

	handler.Handle(w, r)

	results := w.Result()

	expect := 405
	got := results.StatusCode

	if expect != got {
		t.Errorf("Expect %d received %d", expect, got)
	}
}