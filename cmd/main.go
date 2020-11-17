package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/zhubby/blockchain-go"
)

func main() {
	flag.Parse()

	bc := blockchain.NewChain()

	bc.AddBlock("Alice send 1 to Bob")
	bc.AddBlock("Bob send 2 to Alice")

	for _, block := range bc.GetBlocks() {
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
