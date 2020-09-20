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
	Nonce           int64
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) conat() []byte {
	timestamp := strconv.FormatInt(b.Timestamp.Unix(), 10)
	return bytes.Join([][]byte{b.ParentBlockHash, b.Hash, []byte(timestamp)}, []byte{})
}

// set hash sha256(prevhash + data + timestamp)
func (b *Block) SetHash() {
	h := sha256.Sum256(b.conat())
	b.Hash = h[:]
}

// create new block with data and prevhash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	notice, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = int64(notice)
	return block
}
