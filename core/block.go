package core

import (
	"io"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
)

////////////////////////////
// Header									//
////////////////////////////

type Header struct {
	Version 					uint32
	PrevBlockHash			types.Hash
	Timestamp					int64
	Height						uint32
	DataHash 					types.Hash
}


////////////////////////////
// Block									//
////////////////////////////

type Block struct {
	*Header
	Transactions 			[]Transaction
	// Cached version of the header hashed
	hash 							types.Hash

	Validator					crypto.PublicKey
	Signature					*crypto.Signature
}

func NewBlock(h *Header, txs []Transaction) *Block {
	return &Block{
		Header: 			h,
		Transactions: txs,
	}
}

// func (b *Block) Sign()

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if (b.hash.IsZero()) {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}
