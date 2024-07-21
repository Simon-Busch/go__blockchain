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

	if err := b.Verify() ; err != nil {
		return err
	}
	return nil
}
