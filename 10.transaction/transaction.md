# 10. Transaction
## 10.0 Introduction
- `Send Somebody Coins`가 어떤 의미인지 ?
  - **transaction (거래)** 란 무슨 의미인지
  - Unconfirmed / Confirmed는 무슨 의미인지?
  - Mempool이란?
  - Inputs / Outpus ?
- 우리 코인의 회계모델 베이스
  - UTXO (Unspent Transaction Output) - 소비되지 않은 거래 출력값
  - 비트코인도 이 모델을 사용

## 10.1 Introduction to Transactions
- Tx(트랜잭션)에는 Input, Output이 있음
  - 하나의 Input/Output이 있을 수도 있지만 N개가 있을 수도 있음
- **Input**
  - 거래를 실행하기 전에 주머니에 있는 돈
- **Output**
  - 거래 완료 후 각의 사람들이 갖고 있는 액수
  - 거래 이후 돈 재분배한 결과
- Transaction Example
  - When Nico wanna send me $5
    - TxIn[$5(Nico)] : $5있다는것을 증명해야함 
    - TxOut[$5(Me)] : Owner Changed
  - When Nico wanna send me $5, But I only have $10
    - need change $5
    - TxIn[$10(Nico)]
      - it would be reward $10(blockchain)  
    - TxOut[$5(Nico)/$5(Me)]
      - it would be reward $10(miner) --> Coin Base transaction

## 10.2 Coinbase Transaction
- 채굴자들에게 보상을 주기위해 직접 입력을 찍어냄
- 이제 블록을 추가하기 위해 String 입력을 받을 필요 없음

## 10.3 Balances
- `니코`가 블록체인 안에 얼마나 갖고 있는지 확인
  - a. 블록체인 안에 있는 모든 블록체인 찾기
  - b. 각 블록에는 거래내역이 있음
  - c. 거래내역에 있는 `Output`에서 모든 amount를 더해 `Nico`가 얼마나 가지고 있는지 총량을 구함

## 10.4 Mempool
- 거래의 Cycle
 - 거래내역을 보낸다고 바로 블록에 추가되는 것이 아님!
- Mempool
  - 아직 확정되지 않은 (Unconfirmed) 거래내역을 저장함, 아직 어느 블록에 속해있지 않고 Mempool에 저장되어 있음
  - 그렇다면, 언제 확정되는걸까?
    - Miner(=블록을 생성하는 사람)이 블록에 트랜잭션을 저장할 때
    - 거래내역을 확정(Confirm)해 주는 조건으로 수수료(fee)를 가져감
- Confirm
  - 블록에 거래내역을 저장하는 것
  - 1. 블록을 채굴
  - 2. 거래내역을 블록에 저장

## 10.5-6 Add Tx & Make Tx
- Add Transaction : Mempool에 트랜잭션을 추가한다. 
  - 이 때는 From이 중요하지 않음 (지갑으로부터 받을거기 때문에)
  - 그러나 To는 중요하다 (받는 사람)
  ``` Go
  func AddTx(to string, amount int) {
    tx := makeTx("hyewon", to, amount)
    m.Txs = append(m.Txs, tx)
  }
  ```

- Make Transaction : Mempool에 저장될 트랜잭션을 생성
  -	transaction을 생성하게 되면 
    - 보내는 이는 Transaction input을 생성하고
    - 받는 이는 Transaction output을 생성
    - 이 둘을 합쳐서 Transaction을 생성
	- input의 amount와 output의 amount가 같아야 함 
  a. from 사용자의 잔액을 확인 (transaction의 output으로부터 확인하면 됨) 
  b. from 사용자의 output을 가져와서 input으로 사용 
  ``` Go
	oldTxOuts := BlockChain().TxOutsByAddress(from)
  
	for _, txOut := range oldTxOuts {
		// input의 amount가 output의 amount보다 크거나 같아야 함
		if total > amount {
			break
		}
		txIn := &TxIn{txOut.Owner, txOut.Amount}

		// c. input내의 금액을 더해서 sum(txIns.amount) >= 거래금액이 될 때까지 txIns에 추가
		txIns = append(txIns, txIn)
		total += txIn.Amount
	}
  ```
  c. input내의 금액을 더해서 sum(txIns.amount) >= 거래금액이 될 때까지 txIns에 추가 
  d. transaction input금액의 합이 거래금액보다 크다면 잔액을 다시 output으로 생성해주어야 함 -> from 사용자에게 거스름돈을 주는 output을 생성 
  e. to 사용자에게 거래금액을 주는 output을 생성

from : hyewon -> to : lynn (30)
```
[
  {
    "id": "7e1d1eeadc1a1756fc52710d5d7f0c713c0edbf81d20a9ef54de26d8e736e4e8",
    "timestamp": 1733231743,
    "txIns": [
      {
        "Owner": "hyewon",
        "Amount": 50
      }
    ],
    "txOuts": [
      {
        "Owner": "hyewon",
        "Amount": 20
      },
      {
        "Owner": "lynn",
        "Amount": 30
      }
    ]
  }
]
```

## 10.7 Confirm Transactions
- 이제 블록을 생성할 때 마다 mempool에 있는 거래내역을 모두 가져와서 Confirm할 예정
- 블록이 채굴될 때, 거래내역을 확인해서 Confirm
- 각 잔고가 업데이트 되나 문제가 있음!
- Mempool에 들어갈때는 보내는이의 잔고를 비교하지만
- Confirm할때는 검증없이 모두 등록해버림 
  - Confirm할 시점에는 이미 보내는 이의 잔고가 달라질 수도 있는 가능성


**Transaction Confirm**
1. (외부) Transaction 송금 요청
2. Mempool에 Transaction 추가
	1. Transaction 생성 (from : txIn, To : txOut) -> Tx
	2. Mempool에 트랜잭션 Append
3. 블록 생성
	1. 마이닝 이후 Transaction Confirm
    		1. coinbase 트랜잭션 생성 (보상)
    		2. Mempool에 있는 트랜잭션에 Append
    		3. Mempool 비우기


## 10.8 uTxOuts
- 위 코드의 문제점
  - 이 사용자가 이 txOut을 이미 사용했는지 확인하지 않고 Confirm
  - 이미 트랜잭션에 사용된 txOut이 다시 또 Mempool에 넣음
  - => 어떻게 Mempool-tx에 있는 txOut이 이미 사용됐는지 알지?

- TOBE
```
Tx1
  TxIn[COINBASE]
  TxOut[$5(you)] <------ Spent TxOut : 아래 트랜잭션에서 Input으로서 참조했기 때문

Tx2
  TxIn[Tx1.TxOuts[0]] // It should be connected to previous Output *!!!! It'd be a reference to an Old Output 
  TxOut[$5(me)] <---- uTxOut (UnSpent Transaction Output) <----- 아래 Tx3이 생기면서 Spent

Tx3
  TxIns[Tx2.TxOuts[0]] 
  TxOuts[$3(you), $2(me)] <---- uTxOUt X 2 (UTXO) : 잔액 계산시에 보는 곳
```

- 새로운 트랜잭션의 Input은 이전 트랜잭션의 Output을 참조한다
  - 누군가가 트랜잭션 OUTPUT을 실제로 가지고 있는지 여부를 확인할 수 있음


## 10.9 UTxOutsByAddress
- UTxOutsByAddress 함수 구현
  - 아직 input으로부터 사용되지 않은 (참조되지 않은) Output만을 확인
   
## 10.10 makeTx Part 2
- a. from의 잔액이 송금하려는 금액보다 적다면 리턴
- b. UTxOutsByAddress로 UTXO를 가지고 옴
  - 이 UTXO로부터 txIn을 만듬
  - txIn들의 금액을 더한게 송금 amount보다 커지면 break
  - 이 경우 change TxOut도 생성함
- c. TxOut 생성
- 어떤 TxOut이 이미 Mempool에 올라와 있다면 더이상 올라갈수 없어야함 -> 10.11

## 10.11 isOnMempool
- 
