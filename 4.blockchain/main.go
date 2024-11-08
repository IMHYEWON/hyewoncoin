package main

import (
	"fmt"

	"github.com/IMHYEWON/hyewoncoin/4.blockchain/blockchain"
)

func main() {
	// GetBlockChain() 함수를 호출하여 블록체인 인스턴스를 생성
	// 이미 Genesis Block이 생성되어 있음
	chain := blockchain.GetBlockChain()

	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlocks() {
		fmt.Println("Data : ", block.Data)
		fmt.Println("Hash : ", block.Hash)
		fmt.Println("PrevHash : ", block.PrevHash)
	}
}
