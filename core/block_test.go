package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/Simon-Busch/go__blockchain/types"
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
