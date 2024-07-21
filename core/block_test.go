package core

import (
	"fmt"
	"testing"
	"time"
	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header:= &Header{
		Version: 					1,
		PrevBlockHash: 		types.RandomHash(),
		Height: 					height,
		Timestamp: 				time.Now().UnixNano(),
		DataHash: 				types.RandomHash(),
	}

	tx := Transaction{
		Data: []byte("foo bar"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestEncodeDecode(t *testing.T) {
	b := randomBlock(0)
	fmt.Println("Block: ", b)
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestBlockSign(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestBlockVerify(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

	assert.Nil(t, b.Verify())
}
