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
	Transaction  string
	Nonce        int
	Hashptr      *Block
	PreviousHash string
	CurrHash     string
}

func NewBlock(Transaction string, Nonce int, PreviousHash string) *Block {
	b := new(Block)
	/*	b := new(Block{
		Transaction, Nonce, PreviousHash, Hashptr, CalculateHash(fmt.Sprintf("%s%d%s", Transaction, Nonce, PreviousHash)),
	})*/
	b.Transaction = Transaction
	b.Nonce = Nonce
	b.PreviousHash = PreviousHash
	b.Hashptr = nil
	b.CurrHash = CalculateHash(toString(b))
	return b
}
func toString(block *Block) string {
	//var final string
	final := fmt.Sprintf("%s %d %x", block.Transaction, block.Nonce, block.PreviousHash)
	return final
}
func ListBlocks(ptr *Block) {
	iter := ptr
	for iter != nil {
		fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.Transaction, iter.Nonce, iter.CurrHash, iter.PreviousHash)
		iter = iter.Hashptr
	}

}
func ChangeBlock(ptr *Block) {
	iter := ptr
	i := 3
	for i != 0 {
		//fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.Transaction, iter.Nonce, iter.CurrHash, iter.PreviousHash)
		iter = iter.Hashptr
		i--
	}
	iter.Transaction = "B to C"
	iter.CurrHash = CalculateHash(toString(iter))
}
func VerifyChain(ptr *Block) bool {
	iter := ptr
	for iter != nil {
		if iter.PreviousHash != iter.Hashptr.CurrHash {
			fmt.Print("Alert!! Changes were made to the following Transaction\nHashes do not match\n")
			iter = iter.Hashptr
			fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.Transaction, iter.Nonce, iter.CurrHash, iter.PreviousHash)
			return false
		}
		iter = iter.Hashptr
	}
	return true
}
