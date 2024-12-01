package blockchain

import (
	"sync"

	"github.com/IMHYEWON/hyewoncoin/9.mining/db"
	"github.com/IMHYEWON/hyewoncoin/9.mining/utils"
)

const defaultDifficulty = 2
const difficultyInterval = 5

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

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
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

func (b *blockchain) difficulty() int {
	// default difficulty (블로체인이 비어있을 때)
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// 5개의 블록마다 difficulty 재조정
		// recalculate difficulty
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
				bc.AddBlock("Genesis")
			} else {
				bc.restore(checkpoint)
			}
		})
	}

	return bc
}
