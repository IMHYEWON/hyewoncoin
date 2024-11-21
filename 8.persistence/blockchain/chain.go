package blockchain

import (
	"sync"

	"github.com/IMHYEWON/hyewoncoin/8.persistence/db"
	"github.com/IMHYEWON/hyewoncoin/8.persistence/utils"
)

type blockchain struct {
	// blocks []*Block
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"` // 블록의 번호
}

// blockchain.go 내부에서만 이 변수에 접근 가능
var b *blockchain

// sync.Once : 한번만 실행되는 코드를 실행하기 위한 구조체
var once sync.Once

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// 이 메소드로 블록체인 생성을 제어
func BlockChain() *blockchain {
	// nil : 아무것도 없음을 나타내는 특별한 값
	// b가 nil이면 새로운 blockchain을 생성하여 반환
	if b == nil {
		// 블록체인 인스턴스 생성이 단 한번만 실행되도록 보장
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("Genesis")
		})
	}

	// 블록체인 인스턴스 반인
	return b
}
