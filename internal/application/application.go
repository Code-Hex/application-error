package application

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
)

type Service interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key, value string) error
}

func NewGethandler(s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Requires GET", http.StatusMethodNotAllowed)
			return
		}
		key := r.URL.Query().Get("key")
		v, err := s.Get(r.Context(), key)
		if err != nil {
			errorHandling(err, w, r)
			return
		}
		fmt.Fprint(w, v)
	})
}

func NewPuthandler(s Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Requires POST", http.StatusMethodNotAllowed)
			return
		}
		key := r.FormValue("key")
		val := r.FormValue("val")
		log.Printf("key=%s, val=%s", key, val)
		if err := s.Put(r.Context(), key, val); err != nil {
			errorHandling(err, w, r)
			return
		}
		fmt.Fprint(w, "OK")
	})
}

func errorHandling(err error, w http.ResponseWriter, r *http.Request) {
	// 型で比較する
	switch e := err.(type) {
	case *net.OpError:
		// dial tcp [ADDR:PORT]: i/o timeout
		if e.Timeout() {
			http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)
			return
		}
		if e.Temporary() {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
	case *url.Error:
		errorHandling(e.Err, w, r)
		return
	case interface{ Cause() error }: // pkg errors
		errorHandling(e.Cause(), w, r)
		return
	case interface{ Unwrap() error }: // new errors packages
		errorHandling(e.Unwrap(), w, r)
		return
	}

	// 値で比較する
	switch err {
	case context.Canceled, context.DeadlineExceeded:
		http.Error(w, http.StatusText(http.StatusRequestTimeout), http.StatusRequestTimeout)
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
