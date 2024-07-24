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
	RPCHandler 					RPCHandler
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

	s :=  &Server{
		ServerOpts: 				opts,
		memPool: 						NewTxPool(),
		blockTime: 					opts.BlockTime,
		isValidator: 				opts.PrivateKey != nil,
		rpcCh:							make(chan RPC),
		quickCh:						make(chan struct{}, 1),
	}

	if opts.RPCHandler == nil {
		opts.RPCHandler = NewDefaultRPCHandler(s)
	}

	s.ServerOpts = opts

	return s
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)

	free:
		for {
			select {
				case rpc := <- s.rpcCh:
					if err := s.RPCHandler.HandleRPC(rpc); err != nil {
						logrus.Error(err)
					}
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

func (s *Server) ProcessTransaction(from NetAddr, tx *core.Transaction)  error {
	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("Mempool already contains tx")
		return nil
	}

	if err := tx.Verify(); err != nil {
		fmt.Printf("Invalid tx: %s\n", err)
		return err
	}

	tx.SetFirstSeen((time.Now().UnixNano()))

	logrus.WithFields(logrus.Fields{
		"hash": hash,
		"mempool length": s.memPool.Len(),
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
