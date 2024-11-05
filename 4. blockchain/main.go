package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

// Hash function : one way function
// "abc" -> hash("abc") -> "dasdwd342dskd"

func main() {
	genesisBlock := block{"Genesis Block", "", ""}

	// Sum256 function takes a byte slice and returns a byte slice
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	fmt.Println(hash)
	fmt.Printf("%x", hash) // %x -> hexa decimal
	hexHash := fmt.Sprintf("%x", hash)
	genesisBlock.hash = hexHash // get Hash from "Genesis Block" and assign it to hash
	fmt.Println(genesisBlock)

	// 2nd block
	secondBlock := block{"Second Block", "", genesisBlock.hash}

	// Same
	fmt.Println("GenesisBlock Hash: ", hexHash, "Prev Hash: ", secondBlock.prevHash)
	equal := hexHash == secondBlock.prevHash
	fmt.Println("Equal: ", equal)

	secondHash := sha256.Sum256([]byte(secondBlock.data + secondBlock.prevHash))
	secondHexHash := fmt.Sprintf("%x", secondHash)
	secondBlock.hash = secondHexHash
	fmt.Println(secondBlock)
}
