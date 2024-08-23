package core

import (
	"testing"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/stretchr/testify/assert"
)


func TestAccountStateTransferNoBalance(t *testing.T) {
	state := NewAccountState()

	from := crypto.GeneratePrivateKey().PublicKey().Address()
	to := crypto.GeneratePrivateKey().PublicKey().Address()
	amount := uint64(100)

	assert.NotNil(t, state.Transfer(from, to, amount))
}

func TestAccountStateTransferSuccess(t *testing.T) {
	state := NewAccountState()

	from := crypto.GeneratePrivateKey().PublicKey().Address()

	state.AddBalance(from, 1000)

	to := crypto.GeneratePrivateKey().PublicKey().Address()
	amount := uint64(50)

	assert.Nil(t, state.Transfer(from, to, amount))
}
