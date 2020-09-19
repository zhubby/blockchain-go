package main

import (
	"fmt"

	"github.com/zhubby/blockchain-go"
)

func main() {
	bc := blockchain.NewChain()

	bc.AddBlock("Alice send 1 to Bob")
	bc.AddBlock("Bob send 2 to Alice")

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Prev. hash: %x\n", block)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Print("\n")
	}
}