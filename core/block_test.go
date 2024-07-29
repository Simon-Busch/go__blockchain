package core

import (
	"fmt"
	"testing"
	"time"
	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/stretchr/testify/assert"
)


func TestEncodeDecode(t *testing.T) {
	b := randomBlock(t, 0, types.Hash{})
	fmt.Println("Block: ", b)
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestBlockSign(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestBlockVerify(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)

	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)

	header:= &Header{
		Version: 					1,
		PrevBlockHash: 		prevBlockHash,
		Height: 					height,
		Timestamp: 				time.Now().UnixNano(),
		// DataHash: 				types.RandomHash(),
	}

	b, err :=  NewBlock(header, []*Transaction{tx})
	assert.Nil(t, err)
	dataHash, err := CalculateDataHash(b.Transactions)
	assert.Nil(t, err)
	b.Header.DataHash = dataHash
	assert.Nil(t, b.Sign(privKey))

	return b
}
