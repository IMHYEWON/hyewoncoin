package main

import (
	"crypto/sha256"
	"fmt"
)

/***
 * 요구사항
 * - 첫번째 블록을 제외한 모든 블록은 이전 블록의 해시값을 가지고 있어야 한다.
 * - 블록체인에 블록을 추가할 수 있어야 한다.
 * 	chain.addBlock("Genesis block")
	chain.addBlock("Second block")
	chain.addBlock("Third block")
	chain.listBlocks()
***/

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

// 첫번째 블록이 아닌 경우 이전 블록의 해시값을 가져온다.
func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

// 블록을 추가한다.
func (b *blockchain) addBlock(data string) {
	// 이전 블록의 해시값과 데이터를 이용하여 새로운 블록을 생성한다.
	newBlock := block{data, "", b.getLastHash()}

	// 새로운 블록의 해시값을 계산한다.
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)

	// 블록체인에 블록을 추가한다.
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)         // it is a hash of the block sha256(data + prevHash)
		fmt.Printf("PrevHash: %s\n", block.prevHash) // it is a hash of the previous block
	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("Genesis block")
	chain.addBlock("Second block")
	chain.addBlock("Third block")
	chain.listBlocks()
}
