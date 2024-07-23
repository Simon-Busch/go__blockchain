package network

import (
	"fmt"
	"time"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/crypto"
	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports 					[]Transport
	BlockTime 					time.Duration
	PrivateKey 					*crypto.PrivateKey
}

type Server struct {
	ServerOpts
	memPool 						*TxPool
	blockTime 					time.Duration
	isValidator 				bool
	rpcCh 							chan RPC
	quickCh 						chan struct {}
}

func NewServer(opts ServerOpts) *Server {
	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}

	return &Server{
		ServerOpts: 				opts,
		memPool: 						NewTxPool(),
		blockTime: 					opts.BlockTime,
		isValidator: 				opts.PrivateKey != nil,
		rpcCh:							make(chan RPC),
		quickCh:						make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)

	free:
		for {
			select {
				case rpc := <- s.rpcCh:
					fmt.Printf("%-v\n", rpc)
				case <- s.quickCh:
					break free
				case <- ticker.C:
					if s.isValidator {
						s.createNewBlock()
					}
			}
		}
	fmt.Println("Server shutdown")
}

func (s *Server) handleTransaction(tx *core.Transaction)  error {
	if err := tx.Verify(); err != nil {
		fmt.Printf("Invalid tx: %s\n", err)
		return err
	}

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"tx": tx,
			"hash": hash,
		}).Info("Mempool already contains tx")
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"tx": tx,
		"hash": hash,
	}).Info("Adding new TX to the mempool")


	return s.memPool.Add(tx)
}

func (s *Server) createNewBlock() error {
	fmt.Println("Creating a new block")
	return nil
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
