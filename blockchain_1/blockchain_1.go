package main

import (
	"fmt"
	"time"
	"strconv"
	"encoding/hex"
	"crypto/sha256"
)

type Block struct {
	Index int
	Timestamp string
	Data string
	PreviousHash string
	Hash string
}

func calculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) (Block, error) {
	var newBlock Block
	
	t := time.Now()
	
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PreviousHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	
	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
	return false
	}
	
	if oldBlock.Hash != newBlock.PreviousHash {
	return false
	}
	
	if calculateHash(newBlock) != newBlock.Hash {
	return false
	}

	return true
}

func main() {
	genesisBlock := Block{}
	genesisBlock = Block{0, time.Now().String(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	fmt.Printf("%+v\n", genesisBlock)
	
	secondBlock, _ := generateBlock(genesisBlock, "Second Block Data")
	fmt.Printf("%+v\n", secondBlock)
	
	fmt.Println(isBlockValid(secondBlock, genesisBlock))
}