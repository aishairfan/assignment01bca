package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

func CalculateHash(strtoHash string) string { //calculating the sha256 checksum of a string
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strtoHash)))
}

type Block struct {
	/*
		----Structure representing a Block in the chain
		@Transaction - State of transaction
		@Nonce - Randomness 
		@Hashptr - Pointer to the last block
		@PreviousHash - Hash of the previous block
		@CurrHash - Hash of Transaction+Nonce+PreviousHash
	*/
	Transaction  string
	Nonce        int
	Hashptr      *Block
	PreviousHash string
	CurrHash     string
}

func NewBlock(Transaction string, Nonce int, PreviousHash string) *Block { //creating a new block with provided values and returning a pointer
	b := new(Block)
	b.Transaction = Transaction
	b.Nonce = Nonce
	b.PreviousHash = PreviousHash
	b.Hashptr = nil
	b.CurrHash = CalculateHash(toString(b)) //calculating hash of the current block
	return b
}

func toString(block *Block) string { //abstraction for calculating hash of a block
	final := fmt.Sprintf("%s %d %x", block.Transaction, block.Nonce, block.PreviousHash)
	return final
}

func ListBlocks(ptr *Block) { //printing details of each block in a chain gracefully
	iter := ptr
	for iter != nil {
		fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.Transaction, iter.Nonce, iter.CurrHash, iter.PreviousHash)
		iter = iter.Hashptr
	}

}

func ChangeBlock(ptr *Block, transaction string) { //updating transaction and currHash of a block
	ptr.Transaction = transaction
	ptr.CurrHash = CalculateHash(toString(ptr))
}

func VerifyChain(ptr *Block) bool { //verifying the integrity of a chain
	/*
		----Comparison method
		@Compares stored hash of previous block with hash of previous block
	*/
	iter := ptr
	for iter.Hashptr != nil {
		if iter.PreviousHash != iter.Hashptr.CurrHash {
			fmt.Print("\n\n\nAlert!! Changes were made to the following Transaction\nHashes do not match\n")
			iter = iter.Hashptr
			fmt.Printf("\n\nTransaction: %s\n Nonce %d\n Current Hash %s\n Previous Hash %s\n", iter.Transaction, iter.Nonce, iter.CurrHash, iter.PreviousHash)
			return false
		}
		iter = iter.Hashptr
	}
	return true
}
