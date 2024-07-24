package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.(*LocalTransport).peers[trb.Addr()], trb)
	assert.Equal(t, trb.(*LocalTransport).peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Test message")

	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <- trb.Consume()
	buf := make([]byte, len(msg))
	n, err := rpc.Payload.Read(buf)
	assert.Equal(t, n, len(msg))
	assert.Nil(t, err)

	assert.Equal(t, buf, msg)
	assert.Equal(t, rpc.From, tra.Addr())
}
