package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type ChainHash []byte

type Block struct {
	Timestamp       time.Time
	Data            []byte
	ParentBlockHash ChainHash
	Hash            ChainHash
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// set hash sha256(prevhash + data + timestamp)
func (b *Block) SetHash() {
	timestamp := strconv.FormatInt(b.Timestamp.Unix(), 10)
	headers := bytes.Join([][]byte{b.ParentBlockHash, b.Hash, []byte(timestamp)}, []byte{})
	h := sha256.Sum256(headers)
	b.Hash = h[:]
}

// create new block with data and prevhash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
