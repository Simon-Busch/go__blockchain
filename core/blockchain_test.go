package core

import (
	"testing"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
)

func TestSendNativeTransferSuccess(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	signer := crypto.GeneratePrivateKey()
	block := randomBlock(t, uint32(1), getPrevBlockHash(t, bc, uint32(1)))
	assert.Nil(t, block.Sign(signer))

	privKeyBob := crypto.GeneratePrivateKey()
	privKeyAlice := crypto.GeneratePrivateKey()
	addressBob := privKeyBob.PublicKey().Address()
	amount := uint64(100)

	accountBob := bc.accountState.CreateAccount(addressBob)
	accountBob.Balance = amount

	tx := NewTransaction([]byte{})
	tx.From = privKeyBob.PublicKey()
	tx.To = privKeyAlice.PublicKey()
	tx.Value = amount
	tx.Sign(privKeyBob)

	block.AddTransaction(tx)

	assert.Nil(t, bc.AddBlock(block))

	accountAlice, err := bc.accountState.GetAccount(privKeyAlice.PublicKey().Address())
	assert.Nil(t, err)

	assert.Equal(t, accountAlice.Balance, amount)
}


func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlocks := 1000
	for i := 0; i < lenBlocks; i++ {
		block := randomBlock(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)
	assert.NotNil(t, bc.AddBlock(randomBlock(t, 89, types.Hash{})))
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(1))
	assert.False(t, bc.HasBlock(100))
}

func TestGetBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenBlocks := 1000

	for i := 0; i < lenBlocks; i++ {
		block := randomBlock(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
		fetchedBlock, err := bc.GetBlock(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, fetchedBlock, block)
	}
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenBlocks := 1000

	for i := 0; i < lenBlocks; i++ {
		block := randomBlock(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}
}

func TestAddBlockToHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.Nil(t, bc.AddBlock(randomBlock(t, 1, getPrevBlockHash(t, bc, uint32(1)))))
	assert.NotNil(t, bc.AddBlock(randomBlock(t, 3, types.Hash{})))
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(log.NewNopLogger(), randomBlock(t, 0, types.Hash{}))
	assert.Nil(t, err)

	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}
