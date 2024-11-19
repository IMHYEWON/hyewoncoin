package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"` // 블록의 번호
}

type blockchain struct {
	blocks []*Block
}

// blockchain.go 내부에서만 이 변수에 접근 가능
var b *blockchain

// sync.Once : 한번만 실행되는 코드를 실행하기 위한 구조체
var once sync.Once

// 블록의 해시값을 계산해서 리턴
func (bl *Block) getCalculateHash() string {
	hash := sha256.Sum256([]byte(bl.Data + bl.PrevHash))
	bl.Hash = fmt.Sprintf("%x", hash)
	return bl.Hash
}

// 마지막 블록의 해시값을 리턴
func getLastHash() string {
	totalBlocks := len(GetBlockChain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlocks-1].Hash
}

// 새로운 블록을 생성
func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockChain().blocks) + 1}
	newBlock.Hash = newBlock.getCalculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

// 한번만 실행되는 코드
// 이 메소드로 블록체인 생성을 제어
func GetBlockChain() *blockchain {
	// nil : 아무것도 없음을 나타내는 특별한 값
	// b가 nil이면 새로운 blockchain을 생성하여 반환
	if b == nil {
		// 블록체인 인스턴스 생성이 단 한번만 실행되도록 보장
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}

	// 블록체인 인스턴스 반인
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

var ErrNotFound = errors.New("Block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}
