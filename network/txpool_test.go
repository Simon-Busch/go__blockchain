package network

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()

	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAdd(t *testing.T) {
	p := NewTxPool()

	assert.Equal(t, p.Len(), 0)

	tx := core.NewTransaction([]byte("Hello World"))

	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)
}

func TestTxPoolFlush(t *testing.T) {
	p := NewTxPool()

	tx := core.NewTransaction([]byte("Hello World"))
	p.Add(tx)
	assert.Equal(t, p.Len(), 1)

	p.Flush()
	assert.Equal(t, p.Len(), 0)
}

func TestDoubleTxPoolAdd(t *testing.T) {
	p := NewTxPool()

	assert.Equal(t, p.Len(), 0)

	tx := core.NewTransaction([]byte("Hello World"))

	assert.Nil(t, p.Add(tx))
	assert.Equal(t, p.Len(), 1)


	_ = core.NewTransaction([]byte("Hello World"))
	assert.Equal(t, p.Len(), 1)
}

func TestSortTransactions(t *testing.T) {
	p := NewTxPool()
	txLen := 1000

	for i:= 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte(strconv.FormatInt(int64(i), 10))) // All data must be different
		tx.SetFirstSeen(int64(i * rand.Intn(1000)))
		assert.Nil(t, p.Add(tx))
	}

	assert.Equal(t, txLen, p.Len())

	txx := p.Transactions()

	for i := 0; i < len(txx) -1 ; i++ {
		assert.True(t, txx[i].FirstSeen() < txx[i+1].FirstSeen())
	}
}
