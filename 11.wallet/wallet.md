## 11.1 Private and Public Keys
1. We hash the msg
   - `"I love you" -> hash -> "fseefqedsjcnkjnca"`
2. Generate Key Pair|
   - `KeyPair (private Key, Public Key)`
3. Sign the Hash
   - `("fseefqedsjcnkjnca" + private Key) -> 'Signature'`
 
- Then, If somebody asked you to prove that this is your signature, What do you do?
- How are you going to prove that you are the one who signed tha hash.
- -> Verifying by the Public Key!

4. Verify
    - `("fseefqedsjcnkjnca" + 'Signature' + public Key) -> True/false`


## 11.2 Signing Message 

## 11.3 Verifying Message

## 11.4 Restoring Message
- 키 복원 프로세스
    1. Private Key 복구
    2. Private Key로 Public Key 복구
    3. Public Key로 Signature 검증

## 11.12 Transaction Verification
- 내가 새로운 트랜잭션을 생성할 때는, 사용하지 않은 TxOut = UTXO가 있어야 한다. 
- 이 UTXO를 TxIn으로 활용하게 되는데, 이 때 내 거라고 증명하기 위해 TxIn마다 서명을 저장하게 되는 것
- 이 서명은 우리의 private Key로 만들어서 저장할거임
- 이 서명을 검증할 사람들은 Public Key로 증명할 거임, Public Key는 Address이기 때문에 이 UTXO의 Address에 있음
- 그래서 새로운 Transaction의 TxIn을 만들때 TxIn의 부모 UTXO의 address, 만드려는 TxIn의 시그니처와 함께 서명을 검증