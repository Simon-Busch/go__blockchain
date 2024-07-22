package core

import (
	"testing"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/stretchr/testify/assert"
)


func TestSignTransac(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()

	tx := &Transaction {
		Data: []byte("Hello World"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransac(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()

	tx := &Transaction {
		Data: []byte("foo"),
	}

	tx.Sign(privKey)

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()

	tx := &Transaction {
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))

	return tx
}
