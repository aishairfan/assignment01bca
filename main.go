package main

import (
	"math/rand"
	"github.com/aishairfan/assignment01bca"
)

func main() { //driver function
	var tail *assignment01bca.Block //initializing a blockchain

	//adding genesis block
	genesisBlock := assignment01bca.NewBlock("Alice to Bob", 12345, "FirstBlock")
	tail = genesisBlock
	tail.Hashptr = nil
	tail.PreviousHash = "FirstBlock"

	//randomly adding 5 new transactions to the chain
	for i := 0; i < 5; i++ {
		nb := assignment01bca.NewBlock("A to B", int(rand.Float64()), tail.CurrHash)
		nb.Hashptr = tail
		tail = nb
	}

	//adding changeable transaction to the chain
	sadBlock:=assignment01bca.NewBlock("A to B", int(rand.Float64()), tail.CurrHash)
	sadBlock.Hashptr=tail
	tail=sadBlock

	//randomly adding 5 more transaction to the chain
	for i := 0; i < 5; i++ {
		nb := assignment01bca.NewBlock("A to B", int(rand.Float64()), tail.CurrHash)
		nb.Hashptr = tail
		tail = nb
	}

	assignment01bca.ListBlocks(tail) //printing blockchain
	assignment01bca.VerifyChain(tail) //vertfying blockchain before change

	//changing the block data
	assignment01bca.ChangeBlock(sadBlock, "A to C")

	assignment01bca.ListBlocks(tail) //printing blockchain
	assignment01bca.VerifyChain(tail) //vertfying blockchain after change
}
