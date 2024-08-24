package core

import (
	"bytes"
	"encoding/gob"
	"testing"

	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestVerifyTransactionWithTamper(t *testing.T) {
	tx := NewTransaction(nil)

	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()
	hackerPrivKey := crypto.GeneratePrivateKey()

	tx.From = fromPrivKey.PublicKey()
	tx.To = toPrivKey.PublicKey()
	tx.Value = uint64(100)

	assert.Nil(t,tx.Sign(fromPrivKey))

	tx.hash = types.Hash{}

	tx.To = hackerPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func TestNFTTransaction(t *testing.T) {
	collectionTx := CollectionTx{
		Fee: 100,
		MetaData: []byte("foo collection"),
	}

	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		TxInner: 	collectionTx,
	}

	tx.Sign(privKey)
	tx.hash = types.Hash{}

	buf := new(bytes.Buffer)
	assert.Nil(t, gob.NewEncoder(buf).Encode(tx))

	txDecoded := &Transaction{}
	assert.Nil(t, gob.NewDecoder(buf).Decode(txDecoded))
	assert.Equal(t, tx, txDecoded)
}

func TestNativeTransaction(t *testing.T) {
	fromPrivKey := crypto.GeneratePrivateKey()
	toPrivKey := crypto.GeneratePrivateKey()

	tx := &Transaction{
		To: 			toPrivKey.PublicKey(),
		Value: 		100,
	}

	assert.Nil(t, tx.Sign(fromPrivKey))
	assert.NotNil(t, tx.Signature)
}


func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func TestTxEncodeDecode(t *testing.T) {
	tx := randomTxWithSignature(t)
	buf := &bytes.Buffer{}
	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))
	tx.hash = types.Hash{}

	txDecoded := new(Transaction)
	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
	assert.Equal(t, tx, txDecoded)
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("foo"),
	}
	assert.Nil(t, tx.Sign(privKey))

	return &tx
}
