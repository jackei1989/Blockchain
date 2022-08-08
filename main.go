package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash []byte
	Hash     []byte
	Data     []byte
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) SetHash() {
	info := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
} 

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{prevHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return CreateBlock("Genesis Block", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := InitBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	for _, block := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}	
}

