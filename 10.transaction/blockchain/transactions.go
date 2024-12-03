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
	TxId  string `json:"txId"`  // (어떤 transacgion이 이 input을 생성한 output을 가지고 있는지 알려준다)find the previous transaction Output
	Index int    `json:"index"` // UTXO는 같은 (N개이기 때문에) TxId를 가지고 있을 수 있기 때문에 Index로 구분
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// The transaction output that has not been used yet
type UTxOut struct {
	TxId   string `json:"txId"`
	Index  int    `json:"index"`
	Amount int    `json:"amount"`
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

func makeCoinbaseTx(address string) *Tx {
	// 채굴자에게 보상을 주기 위한 트랜잭션
	txIns := []*TxIn{
		{
			TxId:  "",
			Index: -1,
			Owner: "COINBASE",
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

func makeTx(from, to string, amount int) (*Tx, error) {

	if BlockChain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}

	var txIns []*TxIn
	var txOuts []*TxOut

	total := 0

	// from 주소에 해당하는 모든 UTXO를 가져옴
	uTxOuts := BlockChain().UnspentTxOutsByAddress(from)

	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		// 1. 새로운 TxIn 생성
		txIn := &TxIn{
			TxId:  uTxOut.TxId,
			Index: uTxOut.Index,
			Owner: from,
		}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}

	// 2. 새로운 TxOut 생성
	// 2-1. change를 받아야 하는 경우 -> 새로운 TxOut 생성
	// from -> from : change
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{
			Owner:  from,
			Amount: change,
		}
		txOuts = append(txOuts, changeTxOut)
	}

	// 2-2. to에게 보내는 경우 -> 새로운 TxOut 생성
	// from -> to : amount
	txOut := &TxOut{
		Owner:  to,
		Amount: amount,
	}
	txOuts = append(txOuts, txOut)

	// 3. 새로운 Tx 생성
	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil
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

// 이 함수는 블록이 tranasaction을 추가할 때 호출
// 이 때, Mempool에 있는 transaction을 확인하고, Mempool에서 비워줌
// 이 작업은 miner가 새로운 블록을 생성할 때 호출
func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("hyewon")

	// m.Txs : Mempool에 있는 transaction
	txs := m.Txs

	// coinbase transaction을 추가
	txs = append(txs, coinbase)

	// Mempool 비우기
	m.Txs = nil
	return txs
}
