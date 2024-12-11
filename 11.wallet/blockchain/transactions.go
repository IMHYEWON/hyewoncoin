package blockchain

import (
	"errors"
	"time"

	"github.com/IMHYEWON/hyewoncoin/11.wallet/utils"
	"github.com/IMHYEWON/hyewoncoin/11.wallet/wallet"
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
	TxId      string `json:"txId"`  // (어떤 transacgion이 이 input을 생성한 output을 가지고 있는지 알려준다)find the previous transaction Output
	Index     int    `json:"index"` // UTXO는 같은 (N개이기 때문에) TxId를 가지고 있을 수 있기 때문에 Index로 구분
	Signature string `json:"signature"`
}

type TxOut struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
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

func (t *Tx) sign() {
	// 트랜잭션의 모든 TxIn에 대해 서명을 생성해서 저장
	for _, txIn := range t.TxIns {
		txIn.Signature = wallet.Sign(t.Id, wallet.Wallet())
	}
}

// wallet 패키지에 있는 검증 함수를 사용하기 위해서는,
// signature, payload, address 세 가지를 필요로 함
func validate(tx *Tx) bool {
	// 검증하려는 트랜잭션은 새로운 트랜잭션이고, 이 트랜잭션의 txIn은 UTXO로부터 생성된 것이기 때문에
	// UTXO를 가져와서 전송을 시도하려는 이가 진짜로 그 UTXO의 주인인지 확인해야 함 = 서명 확인

	valid := true

	for _, txIn := range tx.TxIns {
		// 만약 여기서 previousTx를 찾을 수 없다면, 이 input은 우리의 블록체인에 존재하지 않는 것
		// = 우리 블록체인 코인을 가지고 있지 않다는 뜻
		previousTx := FindTx(BlockChain(), txIn.TxId)
		if previousTx == nil {
			valid = false
			break
		}

		// 이전 트랜잭션 -> TxOut -> Address -> Public Key 복구
		address := previousTx.TxOuts[txIn.Index].Address
		valid = wallet.Verify(txIn.Signature, tx.Id, address)

	}

	return valid
}

// mempool에 있는 transaction의 input들 중에 uTxOut(트랜잭션에 추가하려는 TxOut 파람)이 있는지 확인
func isOnMempool(uTxOut *UTxOut) bool {
	exists := false

Outer: // label, break Outer로 사용
	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			if input.TxId == uTxOut.TxId && input.Index == uTxOut.Index {
				exists = true
				break Outer
			}
		}
	}

	return exists
}

func makeCoinbaseTx(address string) *Tx {
	// 채굴자에게 보상을 주기 위한 트랜잭션
	txIns := []*TxIn{
		{
			TxId:      "",
			Index:     -1,
			Signature: "COINBASE",
		},
	}

	txOuts := []*TxOut{
		{
			Address: address,
			Amount:  minerReward,
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

// Raw Transaction 생성
func makeTx(from, to string, amount int) (*Tx, error) {

	// from 주소에 해당하는 돈이 충분한지 확인
	if BalanceByAddress(from, BlockChain()) < amount {
		return nil, errors.New("not enough money")
	}

	var txIns []*TxIn
	var txOuts []*TxOut

	total := 0

	// from 주소에 해당하는 모든 UTXO를 가져옴
	uTxOuts := UnspentTxOutsByAddress(from, BlockChain())

	// 목표하는 금액(amount)만큼 UTXO를 찾아서 TxIn에 추가
	// TxIn은 N개의 UTXO를 가질 수 있음
	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		// 1. 새로운 TxIn 생성
		txIn := &TxIn{
			TxId:      uTxOut.TxId,
			Index:     uTxOut.Index,
			Signature: from,
		}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}

	// 2. 새로운 TxOut 생성
	// TxOut은 최대 2개가 될 수 있음
	// 2-1. change를 받아야 하는 경우 -> 새로운 TxOut 생성
	// from -> from : change
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{
			Address: from,
			Amount:  change,
		}
		txOuts = append(txOuts, changeTxOut)
	}

	// 2-2. to에게 보내는 경우 -> 새로운 TxOut 생성
	// from -> to : amount
	txOut := &TxOut{
		Address: to,
		Amount:  amount,
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
	tx.sign()
	return tx, nil
}

// Mempool에 트랜잭션 추가 (트랜잭션을 생성하지는 않음
// 누구에게서 받는지는 중요하지 않음, 누구에게 보내는지만 중요 (지갑으로부터 받을거기 때문)
// REST API 에서 호출할 함수
func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx(wallet.Wallet().Address, to, amount)
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
	coinbase := makeCoinbaseTx(wallet.Wallet().Address)

	// m.Txs : Mempool에 있는 transaction
	txs := m.Txs

	// coinbase transaction을 추가
	txs = append(txs, coinbase)

	// Mempool 비우기
	m.Txs = nil
	return txs
}
