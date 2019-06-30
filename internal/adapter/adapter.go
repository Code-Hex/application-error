package adapter

import "net/http"

// The Adapter type.
// it gets its name from the adapter pattern — also known as the decorator pattern
type Adapter func(http.Handler) http.Handler

func Apply(h http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters) - 1; i >= 0; i-- {
		h = adapters[i](h)
	}
	return h
}
