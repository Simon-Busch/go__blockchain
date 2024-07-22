package core

import (
	"github.com/Simon-Busch/go__blockchain/crypto"
	"fmt"
)

type Transaction struct {
	Data 						[]byte
	From       			crypto.PublicKey
	Signature       *crypto.Signature
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


