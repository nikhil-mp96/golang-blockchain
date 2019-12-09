package main

import (
	"fmt"
	"github.com/nikhil-mp96/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")
	chain.AddBlock("Fourth Block After Genesis")

	for _, v := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", v.PrevHash)
		fmt.Printf("Data: %s \n", v.Data)
		fmt.Printf("Current Block Hash: %x \n", v.Hash)
	}
}
