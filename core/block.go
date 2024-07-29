package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"io"
	"time"

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
	Transactions 				[]*Transaction
	Validator						crypto.PublicKey
	Signature						*crypto.Signature

	// Cached version of the header hashed
	hash 								types.Hash
}

func NewBlock(h *Header, txs []*Transaction) (*Block, error) {
	return &Block{
		Header: 					h,
		Transactions: 		txs,
	}, nil
}

func NewBlockFromPrevHeader(prevHeader *Header, txs []*Transaction) (*Block, error) {
	dataHash, err := CalculateDataHash(txs)
	if err != nil {
		return nil, err
	}

	header := &Header{
		Version: 					prevHeader.Version,
		DataHash: 				dataHash,
		PrevBlockHash: 		BlockHasher{}.Hash(prevHeader),
		Timestamp: 				time.Now().UnixNano(),
		Height: 					prevHeader.Height + 1,
	}
	block, err := NewBlock(header, txs)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, tx)
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

	dataHash, err := CalculateDataHash(b.Transactions)
	if err != nil {
		return err
	}
	if dataHash != b.DataHash {
		return fmt.Errorf("block (%s) has an invalid data hash", b.Hash(BlockHasher{}))
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

func CalculateDataHash(txs []*Transaction) (hash types.Hash, err error) {
	var (
		buf = &bytes.Buffer{}
	)

	for _, tx := range txs {
		if err = tx.Encode(NewGobTxEncoder(buf)) ; err != nil {
			return
		}
	}

	hash = sha256.Sum256(buf.Bytes())

	return
}
