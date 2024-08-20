package api

import (
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/Simon-Busch/go__blockchain/core"
	"github.com/Simon-Busch/go__blockchain/types"
	"github.com/go-kit/log"
	"github.com/labstack/echo/v4"
)

type APIError struct {
	Error string
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
}

type ServerConfig struct {
	Logger        		log.Logger
	ListenAddr				string
	bc 								*core.Blockchain
}

type Server struct {
	ServerConfig
	bc 					*core.Blockchain
}

func NewServer(cfg ServerConfig, bc *core.Blockchain) *Server {
	return &Server{
		ServerConfig: cfg,
		bc:           bc,
	}
}

func (s *Server) Start() error {
	e := echo.New()

	e.GET("/block/:hashorid", s.handleGetBlock)
	return e.Start(s.ListenAddr)
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
	return Block{
		Hash: 					block.Hash(core.BlockHasher{}).String(),
		Version:      	block.Header.Version,
		DataHash:    		block.Header.DataHash.String(),
		PrevBlockHash: 	block.Header.PrevBlockHash.String(),
		Height:      		block.Header.Height,
		Timestamp:   		uint64(block.Header.Timestamp),
		Validator: 	 		block.Validator.Address().String(),
		Signature:   		block.Signature.String(),
	}
}
