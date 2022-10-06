package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

var LastHash string

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Blockhash    string
}

func (b *Block) CalculateHash(stringToHash string) string {

	sum := sha256.Sum256([]byte(stringToHash))
	b.Blockhash = string(sum[:])

	return b.Blockhash

}

type BlockList struct {
	List []*Block
}

func (bc1 *BlockList) NewBlock(Transaction string, Nonce int, PreviousHash string) *Block {

	block := new(Block)
	block.Transaction = Transaction
	block.Nonce = Nonce
	block.PreviousHash = PreviousHash
	block.CalculateHash(Transaction + string(Nonce) + PreviousHash)
	bc1.List = append(bc1.List, block)
	LastHash = block.Blockhash

	return block

}

func ChangeBlock(b1 *Block) {

	b1.Transaction = "Elon Musk to Awais"
	b1.CalculateHash(b1.Transaction + string(b1.Nonce) + b1.PreviousHash)

}

func VerifyChain(ls *BlockList, LastHash string) {

	for i := 0; i < len(ls.List); i++ {

		if i == 0 {
			ls.List[i].CalculateHash(ls.List[i].Transaction + string(ls.List[i].Nonce) + "0")

		} else {

			ls.List[i].CalculateHash(ls.List[i].Transaction + string(ls.List[i].Nonce) + ls.List[i-1].Blockhash)

		}
	}

	if ls.List[len(ls.List)-1].Blockhash == LastHash {

		fmt.Println("\n\n\t\tChain has been Verified")

	} else {
		fmt.Println("\n\n\t\tChain has been CHanged")
	}

}
func (ls *BlockList) Print() {

	for i := range ls.List {

		fmt.Printf("\nTransaction is == %s", ls.List[i].Transaction)

		fmt.Printf("\nNonce is == %d", ls.List[i].Nonce)

		fmt.Printf("\nPrevious hash is == %x", ls.List[i].PreviousHash)

		fmt.Printf("\nBlock hash is == %x", ls.List[i].Blockhash)

	}
}
