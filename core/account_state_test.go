package core

import (
	"testing"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestAccountStateTransferSuccess(t *testing.T) {
	state := NewAccountState()

	address := crypto.GeneratePrivateKey().PublicKey().Address()
	account := state.CreateAccount(address)

	assert.Equal(t, account.Address, address)
	assert.Equal(t, account.Balance, uint64(0))

	fetchedAccount, err := state.GetAccount(address)
	assert.Nil(t, err)
	assert.Equal(t, fetchedAccount, account)

}

func TestTransferFailInsufficientBalance(t *testing.T) {
	state := NewAccountState()

	bobAddress := crypto.GeneratePrivateKey().PublicKey().Address()
	aliceAddress := crypto.GeneratePrivateKey().PublicKey().Address()

	accountBob := state.CreateAccount(bobAddress)
	accountAlice := state.CreateAccount(aliceAddress)

	assert.Equal(t, accountBob.Balance, uint64(0))
	assert.Equal(t, accountAlice.Balance, uint64(0))

	accountBob.Balance = uint64(999)

	amount := uint64(1000)
	assert.ErrorContains(t, state.Transfer(aliceAddress, bobAddress, amount), ErrInsufficientBalance.Error())
}


func TestTransferSuccessEmptyToAccount(t *testing.T) {
	state := NewAccountState()

	bobAddress := crypto.GeneratePrivateKey().PublicKey().Address()
	aliceAddress := crypto.GeneratePrivateKey().PublicKey().Address()

	accountBob := state.CreateAccount(bobAddress)
	accountAlice := state.CreateAccount(aliceAddress)

	assert.Equal(t, accountBob.Balance, uint64(0))
	assert.Equal(t, accountAlice.Balance, uint64(0))

	amount := uint64(1000)

	accountBob.Balance = amount

	assert.Equal(t, accountBob.Balance, amount)
	assert.Nil(t, state.Transfer(bobAddress, aliceAddress, amount))
	assert.Equal(t, accountBob.Balance, uint64(0))
	assert.Equal(t, accountAlice.Balance, amount)
}
