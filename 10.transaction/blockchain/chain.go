package blockchain

import (
	"sync"

	"github.com/IMHYEWON/hyewoncoin/10.transaction/db"
	"github.com/IMHYEWON/hyewoncoin/10.transaction/utils"
)

const defaultDifficulty int = 2  // default difficulty
const difficultyInterval int = 5 // 5개의 블록마다 difficulty 재조정
const blockInterval int = 2      // 2분마다 블록이 생성되도록 설정
const allowedRange int = 2       // 2분의 여유시간
type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`            // 블록의 번호
	CurrentDifficulty int    `json:"currentDifficulty"` // 현재 difficulty
}

var bc *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

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

func (b *blockchain) recalculateDifficulty() int {
	// 5개의 블록마다 시간을 계산 = 5 * 2분 = 10분이 걸렸는지 확인
	allBlocks := b.Blocks()

	// 가장 최근 블록 (allBlocks 함수에서는 가장 최근 hash부터 이전 블록을 차례로 찾아서 추가하기때문에 0번째 인덱스가 가장 최근 블록)
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]

	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval // 5 * 2분 = 10분

	if actualTime < (expectedTime - allowedRange) {
		// 블록이 너무 빠르게 생성되고 있음 (쉬움)> difficulty 높임
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		// 블록이 너무 느리게 생성되고 있음 (어려움)> difficulty 낮춤
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty

}

func (b *blockchain) difficulty() int {
	// default difficulty (블로체인이 비어있을 때)
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// 5개의 블록마다 difficulty 재조정
		// recalculate difficulty
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}

}

func BlockChain() *blockchain {
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{
				Height: 0,
			}
			checkpoint := db.Checkpoint()

			if checkpoint == nil {
				bc.AddBlock()
			} else {
				bc.restore(checkpoint)
			}
		})
	}

	return bc
}
