package health

import (
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	got := NewHandler()

	rec := httptest.NewRecorder()
	got.ServeHTTP(rec, nil)

	if rec.Code != 200 {
		t.Errorf("unexpected %d", rec.Code)
	}
	if rec.Body.String() != "OK" {
		t.Errorf("got %s, want %s", rec.Body.String(), "OK")
	}
}
