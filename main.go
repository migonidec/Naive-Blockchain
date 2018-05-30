package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/hex"
)

type block struct {
	index uint32
	timestamp  time.Time
	data	string
	hash	string
	previousHash string
}

func calculateHash(index uint32, timestamp time.Time, data, previousHash string) string {
	concatString := fmt.Sprint(index) + timestamp.String() + data + previousHash
	hash := sha256.New()
	hash.Write([]byte(concatString))
	return hex.EncodeToString(hash.Sum(nil))
}

func getLastBlock(blockchain []block) block{
	return blockchain[len(blockchain)-1]
}

func addNewBlock(blockchain *[]block, data string) *[]block{
	lastBlock := getLastBlock(*blockchain)
	newBlock := block{
		index: lastBlock.index+1,
		timestamp: time.Now(),
		data: data,
		previousHash: lastBlock.hash,
	}
	newBlock.hash = calculateHash(newBlock.index, newBlock.timestamp, newBlock.data, newBlock.previousHash)

	*blockchain = append(*blockchain, newBlock)

	return blockchain
}


func main(){

	var blockchain []block

	genesisBlock := block{
		index: 0,
		timestamp: time.Now(),
		data: "This is the first naive blockchain block",
		previousHash: "",
	}
	genesisBlock.hash = calculateHash(genesisBlock.index, genesisBlock.timestamp, genesisBlock.data, genesisBlock.previousHash)
	blockchain = append(blockchain, genesisBlock)


	addNewBlock(&blockchain, "hello")

	fmt.Println(blockchain)
}