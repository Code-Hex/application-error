package storage

import (
	"context"
	"errors"
)

var ErrNotFoundItem = errors.New("item not found")

type Service interface {
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key, value string) error
}
