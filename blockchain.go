package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	prevHash     string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, prevHash string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.prevHash = prevHash
	b.timestamp = time.Now().UnixNano()
	return b

}

func (b *Block) Print() {
	fmt.Printf("timestamp:			%d\n", b.timestamp)
	fmt.Printf("prevHash:			%s\n", b.prevHash)
	fmt.Printf("nonce:				%d\n", b.nonce)
	fmt.Printf("transactions:			%s\n", b.transactions)

}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.InsertBlockInBlockChain(0, "init hash")
	return bc
}

func (bc *Blockchain) InsertBlockInBlockChain(nonce int, prevHash string) {
	b := NewBlock(nonce, prevHash)
	bc.chain = append(bc.chain, b)
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s \n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain:	")
}
func main() {
	bc := NewBlockChain()
	bc.Print()
	bc.InsertBlockInBlockChain(4, "hash1")
	bc.Print()
	bc.InsertBlockInBlockChain(5, "hash2")
	bc.Print()

}
