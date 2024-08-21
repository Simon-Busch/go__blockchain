package api

import (
	"encoding/gob"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/go-kit/log"
	"github.com/labstack/echo/v4"
)

type TxResponse struct {
	TxCount 				uint
	Hashes					[]string
}

type APIError struct {
	Error 					string
}

type Block struct {
	Version 		 		uint32
	Hash 						string
	DataHash 		 		string
	PrevBlockHash 	string
	Height 					uint32
	Timestamp 			uint64
	Validator 			string
	Signature 			string

	TxResponse 			TxResponse
}

type ServerConfig struct {
	Logger        		log.Logger
	ListenAddr				string
}

type Server struct {
	txChan 					chan *core.Transaction
	ServerConfig
	bc 								*core.Blockchain
}

func NewServer(cfg ServerConfig, bc *core.Blockchain, txChan chan *core.Transaction) *Server {
	return &Server{
		ServerConfig:  cfg,
		bc:          	 bc,
		txChan:        txChan,
	}
}

func (s *Server) Start() error {
	e := echo.New()

	e.GET("/block/:hashorid", s.handleGetBlock)
	e.GET("/tx/:hash", s.handleGetTX)
	e.POST("/tx", s.HandlePostTx)

	return e.Start(s.ListenAddr)
}

func (s *Server) HandlePostTx(c echo.Context) error {
	tx := &core.Transaction{}

	if err := gob.NewDecoder(c.Request().Body).Decode(tx); err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}

	s.txChan <- tx

	return nil
}

func (s *Server) handleGetTX(c echo.Context) error {
	hash := c.Param("hash")

	b, err := hex.DecodeString(hash)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}
	h := types.HashFromBytes(b)

	tx, err := s.bc.GetTxByHash(h)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, tx)
}

func (s *Server) handleGetBlock(c echo.Context) error {
	idOrHash := c.Param("hashorid")

	height, err := strconv.Atoi(idOrHash)

	// If err is nil we can assume the height of the block is passed
	if err == nil {
		block, err := s.bc.GetBlock(uint32(height))

		if err != nil {
			// return c.JSON(http.StatusBadRequest, map[string]any{"error": err})
			return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, intoJSONBlock(block))
	}
	// Otherwise we assume the hash of the block is passed

	b, err := hex.DecodeString(idOrHash)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}
	h := types.HashFromBytes(b)

	block, err := s.bc.GetBlockByHash(h)
	if err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, intoJSONBlock(block))
}

func intoJSONBlock(block *core.Block) Block {
	txResponse := TxResponse{
		TxCount: uint(len(block.Transactions)),
		Hashes:  make([]string, len(block.Transactions)),
	}

	for i := 0 ; i < int(txResponse.TxCount) ; i++ {
		txResponse.Hashes[i] = block.Transactions[i].Hash(core.TxHasher{}).String()
	}

	return Block{
		Hash: 					block.Hash(core.BlockHasher{}).String(),
		Version:      	block.Header.Version,
		DataHash:    		block.Header.DataHash.String(),
		PrevBlockHash: 	block.Header.PrevBlockHash.String(),
		Height:      		block.Header.Height,
		Timestamp:   		uint64(block.Header.Timestamp),
		Validator: 	 		block.Validator.Address().String(),
		Signature:   		block.Signature.String(),
		TxResponse: 		txResponse,
	}
}
