package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"encoding/hex"
	"encoding/json"
	"crypto/sha256"
)

type Block struct {
	Index int
	Data map[string]interface{}
	PreviousHash string
	Hash string
	Timestamp time.Time
	ProofOfWork int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockInfo := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.ProofOfWork)
	blockHash := sha256.Sum256([]byte(blockInfo))
	blockHashStr := hex.EncodeToString(blockHash[:])
	return fmt.Sprintf("%s", blockHashStr)
} 

func (b *Block) mine(minimumEffort int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", minimumEffort)) {
		b.ProofOfWork++
		b.Hash = b.calculateHash()
	}
	fmt.Printf("Hash: %s | POW: %d\n", b.Hash, b.ProofOfWork)
}

