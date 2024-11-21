package blockchain

import (
	"fmt"
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

func (b *blockchain) restore(data []byte) {
	// blockchain 구조체에 저장된 데이터를 디코딩
	// argument should be a pointer to the value
	// decode will replace the value on the memory address
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

// 모든 블록을 가져오는 함수
// Newest Block으로부터 prevHash를 따라가면서 모든 블록을 가져옴
func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash

	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)

		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

// 이 메소드로 블록체인 생성을 제어
func BlockChain() *blockchain {
	// nil : 아무것도 없음을 나타내는 특별한 값
	// b가 nil이면 새로운 blockchain을 생성하여 반환
	if b == nil {
		// 블록체인 인스턴스 생성이 단 한번만 실행되도록 보장
		once.Do(func() {
			b = &blockchain{"", 0}

			// it would be nothing
			fmt.Printf("NewestHash: %s, Height: %d\n", b.NewestHash, b.Height)

			// DB에 저장된 checkpoint를 가져와서 블록체인에 반영
			// DB에는 Byte로 저장되어 있으므로 Byte를 다시 blockchain으로 변환
			checkpoint := db.Checkpoint()

			if checkpoint == nil {
				// Genesis Block 생성
				b.AddBlock("Genesis")
			} else {
				// restore the blockchain (& decode)
				b.restore(checkpoint)
				fmt.Printf("Restored data: %v\n", b)
			}
		})
	}

	fmt.Printf("NewestHash: %s, Height: %d\n", b.NewestHash, b.Height)
	// 블록체인 인스턴스 반인
	return b
}
