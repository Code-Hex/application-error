package handler

import (
	"net/http"

	"github.com/Code-Hex/application-error/internal/adapter"
)

type Handler struct {
	mux      *http.ServeMux
	adapters []adapter.Adapter
}

func New(adapters ...adapter.Adapter) *Handler {
	return &Handler{
		mux:      http.NewServeMux(),
		adapters: adapters,
	}
}

func (h *Handler) Handle(path string, hh http.Handler) {
	h.mux.Handle(path, adapter.Apply(hh, h.adapters...))
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// server 構造体が持つ mux のハンドラーへリクエストを流す
	h.mux.ServeHTTP(w, r)
}
