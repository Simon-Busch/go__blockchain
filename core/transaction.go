package core

import (
	"fmt"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
)

type Transaction struct {
	Data 						[]byte
	From       			crypto.PublicKey
	Signature       *crypto.Signature

	// Cached version of the tx data hash
	hash 						types.Hash
	// Timestamp of when the tx was first seen locally

	firstSeen 			uint64
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}

	return tx.hash
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.Signature = sig
	tx.From = privKey.PublicKey()

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("tx has no signature")
	}
	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("tx signature is invalid")
	}

	return nil
}

func (tx *Transaction) SetFirstSeen(t int64) {
	tx.firstSeen = uint64(t)
}

func (tx *Transaction) FirstSeen() uint64 {
	return tx.firstSeen
}
