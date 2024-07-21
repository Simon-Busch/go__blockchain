package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	return bc
}

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
		b := randomBlockWithSignature(t, uint32(i + 1))
		assert.Nil(t, bc.AddBlock(b))
		assert.Equal(t, bc.Height(), uint32(i + 1))
	}

	assert.Equal(t, bc.Height(), uint32(1000))
	assert.Equal(t, len(bc.headers), 1001)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 89))) // Will throw as it already exists
}
