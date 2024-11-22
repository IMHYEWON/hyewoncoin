package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/IMHYEWON/hyewoncoin/9.mining/db"
	"github.com/IMHYEWON/hyewoncoin/9.mining/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"` // 블록의 번호
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}

	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}

	payload := block.Data + block.PrevHash + string(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))

	block.persist()

	return block
}
