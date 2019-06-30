package storage

import (
	"context"
	"net"
	"sync"
)

type timeout struct{}

func (*timeout) Timeout() bool {
	return true
}

func (*timeout) Error() string { return "timeout" }

const timeoutKey = "timeout"

var errTimeout = &net.OpError{Err: &timeout{}}

// InMemStore is a store backed by a map that should only be used for testing.
type InMemStore struct {
	sync.RWMutex
	m map[string]string
}

func (i *InMemStore) Get(ctx context.Context, key string) (string, error) {
	i.RLock()
	defer i.RUnlock()
	if key == timeoutKey {
		return "", errTimeout
	}
	v, ok := i.m[key]
	if !ok {
		return "", ErrNotFoundItem
	}
	return v, nil
}

func (i *InMemStore) Put(ctx context.Context, key, value string) error {
	i.Lock()
	if key == timeoutKey {
		return errTimeout
	}
	if i.m == nil {
		i.m = map[string]string{}
	}
	i.m[key] = value
	i.Unlock()
	return nil
}
