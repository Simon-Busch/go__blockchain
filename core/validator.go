package core

import "fmt"


type Validator interface {
	ValidateBlock(b *Block) error
}

type BlockValidator struct {
	bc 								*Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (bv *BlockValidator) ValidateBlock(b *Block) error {
	if bv.bc.HasBlock(b.Height) {
		return fmt.Errorf("block with height %d already exists with hash (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height != bv.bc.Height() + 1 {
		return fmt.Errorf("block height is invalid, expected %d, got %d", bv.bc.Height() + 1, b.Height)
	}

	prevHeader, err := bv.bc.GetHeader(b.Height - 1)
	if err != nil {
		return fmt.Errorf("could not get previous header: %s", err)
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if hash != b.PrevBlockHash {
		return fmt.Errorf("block prev hash is invalid, expected %s, got %s", hash, b.PrevBlockHash)
	}

	if err := b.Verify() ; err != nil {
		return err
	}

	return nil
}
