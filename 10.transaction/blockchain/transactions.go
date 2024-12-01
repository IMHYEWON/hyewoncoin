package blockchain

import (
	"errors"
	"time"

	"github.com/IMHYEWON/hyewoncoin/10.transaction/utils"
)

const minerReward int = 50

type mempool struct {
	Txs []*Tx
}

// Mempool is on the memory, so it is not saved in the database
// If transaction is confirmed, it will be deleted from the mempool and added to the block
var Mempool *mempool = &mempool{}

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

// 내 transaction output : [1, 1, 1, 1, 1, 1, 1, 1, 1, 1]
// 다른 이에게 $5를 보내고 싶다면?
// transaction input : [1, 1, 1, 1, 1]을 위 output으로부터 가져와야 함

func makeTx(from, to string, amount int) (*Tx, error) {
	// transaction을 생성하게 되면
	// 보내는 이는 Transaction input을 생성하고
	// 받는 이는 Transaction output을 생성
	// 이 둘을 합쳐서 Transaction을 생성
	// inout의 amount와 output의 amount가 같아야 함

	// a. from 사용자의 잔액을 확인 (transaction의 output으로부터 확인하면 됨)
	if BlockChain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}

}

// Mempool에 트랜잭션 추가 (트랜잭션을 생성하지는 않음
// 누구에게서 받는지는 중요하지 않음, 누구에게 보내는지만 중요 (지갑으로부터 받을거기 때문)
// REST API 에서 호출할 함수
func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("hyewon", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}
