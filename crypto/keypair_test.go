package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	address := pubKey.Address()

	assert.NotNil(t, privKey)
	assert.NotNil(t, pubKey)

	fmt.Println("Address: ",address)

	msg := []byte("Hello World")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	fmt.Println("Signature: ", sig)

	assert.True(t, sig.Verify(pubKey, msg))

	wrongMessage := []byte("Hello!")
	assert.False(t, sig.Verify(pubKey, wrongMessage))
}
