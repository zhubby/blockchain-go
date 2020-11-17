package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
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

func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	return result.Bytes(), err
}

func (b *Block) Deserialize(bi []byte) error {
	decoder := gob.NewDecoder(bytes.NewReader(bi))
	return decoder.Decode(b)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) conat() []byte {
	timestamp := strconv.FormatInt(b.Timestamp.Unix(), 10)
	return bytes.Join([][]byte{b.ParentBlockHash, b.Hash, []byte(timestamp)}, []byte{})
}

// set hash sha256(prevHash + data + timestamp)
func (b *Block) SetHash() {
	h := sha256.Sum256(b.conat())
	b.Hash = h[:]
}

// create new block with data and prevHash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	notice, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = int64(notice)
	return block
}
