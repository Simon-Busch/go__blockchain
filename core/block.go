package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)
	return buf.Bytes()
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
		Header: 					h,
		Transactions: 		txs,
	}
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, *tx)
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.Header.Bytes())

	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("block signature is invalid")
	}

	for _, tx := range b.Transactions {
		if err := tx.Verify() ; err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(b)
}

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if (b.hash.IsZero()) {
		b.hash = hasher.Hash(b.Header)
	}

	return b.hash
}
