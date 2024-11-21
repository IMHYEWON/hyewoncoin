package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/IMHYEWON/hyewoncoin/8.persistence/db"
	"github.com/IMHYEWON/hyewoncoin/8.persistence/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"` // 블록의 번호
}

func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	err := encoder.Encode(b)

	utils.HandleErr(err)

	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}

	// Data, PrevHash, Height를 합쳐서 Hash를 생성
	payload := block.Data + block.PrevHash + string(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))

	block.persist()

	return block
}
