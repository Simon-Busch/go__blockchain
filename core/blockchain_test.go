package core

import (
	"fmt"
	"testing"

	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	bc:= newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println("Height: ",bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc:= newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(1))
}

func TestAddBlock(t *testing.T) {
	bc:= newBlockchainWithGenesis(t)
	for i := 0 ; i < 1000 ; i ++ {
		b := randomBlockWithSignature(t, uint32(i + 1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
		assert.Equal(t, bc.Height(), uint32(i + 1))
	}

	assert.Equal(t, bc.Height(), uint32(1000))
	assert.Equal(t, len(bc.headers), 1001)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 89, types.Hash{}))) // Will throw as it already exists
}


func TestAddBlockTooHigh(t *testing.T) {
	bc:= newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3, types.Hash{})))
}

func TestGetHeader(t *testing.T) {
	bc:= newBlockchainWithGenesis(t)
	lenBlocks := 1000
	for i := 0 ; i < lenBlocks ; i ++ {
		b := randomBlockWithSignature(t, uint32(i + 1), getPrevBlockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))

		header, err := bc.GetHeader(b.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, b.Header)
	}
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)
	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevHeader)
}
