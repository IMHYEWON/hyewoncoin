package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IMHYEWON/hyewoncoin/10.transaction/db"
	"github.com/IMHYEWON/hyewoncoin/10.transaction/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
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
	target := strings.Repeat("0", b.Difficulty)

	for {
		// 언제 블록을 생성했는지 확언
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n[%s]\n Block as String: %s\nTarget:%s\n Hash: %s\nNonce: %d\n\n\n", fmt.Sprint(b.Height), fmt.Sprint(b), target, hash, b.Nonce)

		// hash값이 target값으로 시작하는지 확인
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: BlockChain().difficulty(),
		Nonce:      0,
	}

	// 블록을 마이닝하고 저장
	block.mine()
	block.Transactions = Mempool.TxToConfirm()
	block.persist()

	return block
}
