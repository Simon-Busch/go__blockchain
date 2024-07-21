package core


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
	// Validate the block first
	return nil
	
}
