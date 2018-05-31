package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type block struct {
	index uint32
	timestamp  time.Time
	data	string
	hash	string
	previousHash string
}

func calculateHash(index uint32, timestamp time.Time, data, previousHash string) string {
	record := fmt.Sprint(index) + timestamp.String() + data + previousHash
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

func getLastBlock(blockchain []block) block{
	return blockchain[len(blockchain)-1]
}

func isNewBlockValid(blockchain *[]block, newBLock block) bool{
	lastBlock := getLastBlock(*blockchain)
	if (strings.Compare(lastBlock.hash, newBLock.previousHash) != 0)&&(lastBlock.index+1 == newBLock.index) { return false }
	return true
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

	if isNewBlockValid(blockchain, newBlock) { *blockchain = append(*blockchain, newBlock) }

	return blockchain
}

func printBlockchain(blockchain []block){
	fmt.Printf("INDEX\t\tDATA\t\t\t\tTIMESTAMP\n")
	for i:=0; i<len(blockchain); i++ {
		currentBlock := blockchain[i]
		fmt.Printf("%d\t\t%s\t\t\t%s\n", currentBlock.index, currentBlock.data, currentBlock.timestamp.Format("02/01/2006 15:04:05.99"))
	}
}

func main(){

	var blockchain []block

	genesisBlock := block{
		index: 0,
		timestamp: time.Now(),
		data: "first block",
		previousHash: "",
	}
	genesisBlock.hash = calculateHash(genesisBlock.index, genesisBlock.timestamp, genesisBlock.data, genesisBlock.previousHash)
	blockchain = append(blockchain, genesisBlock)

	addNewBlock(&blockchain, "hello there")
	printBlockchain(blockchain)

}