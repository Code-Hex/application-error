package application

import (
	"net"
	"testing"
)

type timeout struct{}

func (*timeout) Timeout() bool {
	return true
}

func (*timeout) Error() string { return "timeout" }

var errTimeout = &net.OpError{Err: &timeout{}}

func TestHello(t *testing.T) {
	_ = errTimeout
}
