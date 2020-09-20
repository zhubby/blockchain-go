package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

const targetBits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func (p *ProofOfWork) prepareData(nonce int64) []byte {
	return bytes.Join([][]byte{
		p.block.ParentBlockHash,
		p.block.Data,
		[]byte(strconv.FormatInt(p.block.Timestamp.Unix(), 16)),
		[]byte(strconv.FormatInt(int64(targetBits), 16)),
		[]byte(strconv.FormatInt(nonce, 16)),
	}, []byte{})
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (p ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Printf("Mining the block containing \"%s\"\n", p.block.Data)
	for nonce < math.MaxInt64 {
		data := p.prepareData(int64(nonce))
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(p.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(p.target) == -1

	return isValid
}
