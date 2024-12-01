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