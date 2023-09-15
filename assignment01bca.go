package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

func CalculateHash(strtoHash string) string {
	//fmt.Printf("String Received:     %s\n", strtoHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strtoHash)))
}

type Block struct {
	transaction  string
	nonce        int
	hashptr      *Block
	previousHash string
	currHash     string
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	b := new(Block)
	/*	b := new(Block{
		transaction, nonce, previousHash, hashptr, CalculateHash(fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)),
	})*/
	b.transaction = transaction
	b.nonce = nonce
	b.previousHash = previousHash
	b.hashptr = nil
	b.currHash = CalculateHash(toString(b))
	return b
}
func toString(block *Block) string {
	//var final string
	final := fmt.Sprintf("%s %d %x", block.transaction, block.nonce, block.previousHash)
	return final
}
func ListBlocks(ptr *Block) {
	iter := ptr
	for iter != nil {
		fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.transaction, iter.nonce, iter.currHash, iter.previousHash)
		iter = iter.hashptr
	}

}
func ChangeBlock(ptr *Block) {
	iter := ptr
	i := 3
	for i != 0 {
		//fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.transaction, iter.nonce, iter.currHash, iter.previousHash)
		iter = iter.hashptr
		i--
	}
	iter.transaction = "B to C"
	iter.currHash = CalculateHash(toString(iter))
}
func VerifyChain(ptr *Block) bool {
	iter := ptr
	for iter != nil {
		if iter.previousHash != iter.hashptr.currHash {
			fmt.Print("Alert!! Changes were made to the following transaction\nHashes do not match\n")
			iter = iter.hashptr
			fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.transaction, iter.nonce, iter.currHash, iter.previousHash)
			return false
		}
		iter = iter.hashptr
	}
	return true
}
