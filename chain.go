package blockchain

type Blocks []*Block
type Chain struct {
	blocks Blocks
}

// add new block into chain
func (c *Chain) AddBlock(data string) {
	prevBlock := c.blocks[len(c.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, newBlock)
}

func (c *Chain) GetBlocks () Blocks{
	return c.blocks
}

func NewChain() *Chain{
	return &Chain{[]*Block{NewGenesisBlock()}}
}