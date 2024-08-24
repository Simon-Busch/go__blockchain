package core

import (
	"encoding/gob"
	"fmt"
	"math/rand"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
)

type TxType byte

const (
	TxTypeCollection TxType = iota // will increment below values by 1 -- 0x0
	TxTypeMint // 0x1
)

type CollectionTx struct {
	Fee 					int64
	MetaData 			[]byte
}

type MintTx struct {
	Fee 							int64
	NFT 							types.Hash
	Collection 				types.Hash
	MetaData 					[]byte // Could be a jpg or wathever
	CollectionOwner 	crypto.PublicKey
	Signature 				crypto.Signature
}

type Transaction struct {
	// Only used for native NFT logic
	TxInner 			any // TxInner will ultimately be either the CollectionTx or the MintTx
	Data      		[]byte
	To 						crypto.PublicKey
	Value 				uint64 // Should be big.Int ultimately
	From      		crypto.PublicKey
	Signature 		*crypto.Signature
	Nonce 				int64

	// cached version of the tx data hash
	hash 					types.Hash
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
		Nonce: rand.Int63n(10000000000000000),
	}
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}
	return tx.hash
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	hash := tx.Hash(&TxHasher{})
	sig, err := privKey.Sign(hash.ToSlice())
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	hash := tx.Hash(&TxHasher{})

	if !tx.Signature.Verify(tx.From, hash.ToSlice()) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

// Due to the fact that TxInner 			any we need to register the types
func init() {
	gob.Register(CollectionTx{})
	gob.Register(MintTx{})
}
