package network

import (
	"sort"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/types"
)

type TxPool struct {
	transactions  					map[types.Hash]*core.Transaction
}

type TxMapSorter struct {
	transactions 										[]*core.Transaction
}

func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	transactions := make([]*core.Transaction, len(txMap))

	i := 0
	for _, tx := range txMap {
		transactions[i] = tx
		i++
	}

	s := &TxMapSorter{}

	sort.Sort(s)
	return s
}

func (s *TxMapSorter) Len() int {
	return len(s.transactions)
}

func (s *TxMapSorter) Swap( i, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}

func (s *TxMapSorter) Less(i, j int) bool {
	return s.transactions[i].FirstSeen() < s.transactions[j].FirstSeen()
}

func (p *TxPool) Transactions() []*core.Transaction {
	s := NewTxMapSorter(p.transactions)
	return s.transactions
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

// Add adds a Tx to the pool, the caller is reponsible for checking if the Tx is already in the pool
func (tp *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})

	// if tp.Has(hash) {
	// 	return nil
	// }

	tp.transactions[hash] = tx

	return nil
}

func (tp *TxPool) Has(hash types.Hash) bool {
	_, ok := tp.transactions[hash]
	return ok
}
