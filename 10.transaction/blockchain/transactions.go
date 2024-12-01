package blockchain

import (
	"time"

	"github.com/IMHYEWON/hyewoncoin/10.transaction/utils"
)

const minerReward int = 50

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txIns"`
	TxOuts    []*TxOut `json:"txOuts"`
}

type TxIn struct {
	Owner  string
	Amount int
}

type TxOut struct {
	Owner  string
	Amount int
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func makeCoinbaseTx(address string) *Tx {
	// 채굴자에게 보상을 주기 위한 트랜잭션
	txIns := []*TxIn{
		{
			Owner:  "COINBASE",
			Amount: minerReward,
		},
	}

	txOuts := []*TxOut{
		{
			Owner:  address,
			Amount: minerReward,
		},
	}

	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}

	tx.getId()
	return &tx
}
