package health

import (
	"fmt"
	"net/http"
)

// NewHandler returns handler to check health of this application.
func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})
}
