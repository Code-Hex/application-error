package adapter

import (
	"bytes"
	"net/http"
	"testing"
)

func TestApply(t *testing.T) {
	var buf bytes.Buffer
	h := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		buf.WriteString("main")
	})

	got := Apply(h,
		func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf.WriteString("A")
				h.ServeHTTP(w, r)
				buf.WriteString("A")
			})
		},
		func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf.WriteString("B")
				h.ServeHTTP(w, r)
				buf.WriteString("B")
			})
		},
		func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				buf.WriteString("C")
				h.ServeHTTP(w, r)
				buf.WriteString("C")
			})
		},
	)
	got.ServeHTTP(nil, nil)
	if buf.String() != "ABCmainCBA" {
		t.Fatalf("got %v, but want `ABCmainCBA`", buf.String())
	}
}
