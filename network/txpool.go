package network

import (
	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/types"
)

type TxPool struct {
	transactions  map[types.Hash]*core.Transaction
}


func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (tp *TxPool) Len() int {
	return len(tp.transactions)
}

func (tp *TxPool) Flush() {
	tp.transactions = make(map[types.Hash]*core.Transaction)
}

func (tp *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})

	if tp.Has(hash) {
		return nil
	}

	tp.transactions[hash] = tx

	return nil
}

func (tp *TxPool) Has(hash types.Hash) bool {
	_, ok := tp.transactions[hash]
	return ok
}
