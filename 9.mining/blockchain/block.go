package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"

	"github.com/IMHYEWON/hyewoncoin/9.mining/db"
	"github.com/IMHYEWON/hyewoncoin/9.mining/utils"
)

// 우리의 블록체인의 작업증명 규칙
// 0이 2개로 시작하는 해시를 찾는다
const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
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

func (b *Block) mine() {
	target := strings.Repeat("0", difficulty)

	for {
		blockAsString := fmt.Sprint(b)

		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("Block as String: %s\nHash: %s\nNonce: %d\n\n", blockAsString, hash, b.Nonce)

		// hash값이 target값으로 시작하는지 확인
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}

	// 블록을 마이닝하고 저장
	block.mine()
	block.persist()

	return block
}
