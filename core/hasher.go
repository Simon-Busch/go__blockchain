package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/Simon-Busch/go__blockchain/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Header) types.Hash {
	h := sha256.Sum256(b.Bytes())
	return types.Hash(h)
}

type TxHasher struct{}

// Bytes
// data ?
// to 32
// value 8
// from 32
// nonce 8
func (TxHasher) Hash(tx *Transaction) types.Hash {
	buf := new(bytes.Buffer)
	if err := gob.NewEncoder(buf).Encode(tx); err != nil {
		panic(err)
	}

	return types.Hash(sha256.Sum256(buf.Bytes()))
}
