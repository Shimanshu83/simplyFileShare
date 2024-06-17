package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcsTransport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)
}
