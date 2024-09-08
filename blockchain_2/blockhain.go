package main

import (
	"time"
)

type Blockchain struct {
	GenesisBlock Block
	Chain []Block
	MinimunEffort int // How computationally expensive is to mine a block or how many zeros the hash must start with
}

func (bc *Blockchain) addBlock(from, to string, amount float64) {
	newBlockData := map[string]interface{}{
		"from": from,
		"to": to,
		"amount": amount,
	}

	lastBlock := bc.Chain[len(bc.Chain)-1]

	newBlock := Block{
		Index: lastBlock.Index + 1,
		Data: newBlockData,
		PreviousHash: lastBlock.Hash,
		Timestamp: time.Now(),
	}

	newBlock.mine(bc.MinimunEffort)

	bc.Chain = append(bc.Chain, newBlock)
}

func (bc Blockchain) isValid() bool {
	for i := range bc.Chain[1:] {
		previousBlock := bc.Chain[i]
		currentBlock := bc.Chain[i+1]
		if currentBlock.Hash != currentBlock.calculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}

	return true
}

func generateBlockchain(minimumEffort int) Blockchain {
	genesisBlock := Block{
		Hash: "0",
		Timestamp: time.Now(),
		Data: map[string]interface{}{
			"info": "Genesis Block",
		},
		Index: 0,
	}

	blockchain := Blockchain{
		GenesisBlock: genesisBlock,
		Chain: []Block{genesisBlock},
		MinimunEffort: minimumEffort,
	}

	return blockchain
}

