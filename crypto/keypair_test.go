package crypto

import (
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.NotNil(t, privKey)

	pubKey := privKey.PublicKey()
	assert.NotNil(t, pubKey)

	// address := pubKey.Address()

}
