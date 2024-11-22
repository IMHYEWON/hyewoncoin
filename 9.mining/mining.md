# 9.0 Introduction to POW
- 작업증명을 구현해서, 채굴자들이 내 블록체인에 블록을 추가할 수 있도록
  - 작업증명을 구현하면, 블록을 추가하기가 어려워짐 -> 이 작업을 해서 블록을 추가하는 이들에게 보상을지급
  - 비트코인, 이더리음 같은 암호화폐에서 사용
    - 최근의 암호화폐와 같은 경우는, 지분증명(PoS)를 이용
- 블록을 추가할 때, 채굴자는 해당 블록에 있는 모든 transaction들을 검증하고 보상을 받음.


# 9.1 PoW Proof Of Work Concept
- 작업증명을 구현
  - `어떻게 해서` 컴퓨터가 블록을 추가하기 어렵게 할 수 있을까?
  - `컴퓨터가 풀기는 어렵지만 검증하기는 쉬운` 방법이 필요
- **우리 프로그램의 POW**
  - 우리 프로그램의 작업 증명 : **n개의 0으로 시작하는 hash를 찾는 것**
    - Hash : One-way function 이기 때문에 출력값을 예측할 수 없음
    - Difficulty : 시작 부분에 몇개의 0이 될지는 블록의 Difficulty에 의해서 결정
    - Nonce : 채굴자들이 유일하게 바꿀 수 있는 값
 - 개념증명 (Proof Of Concept)
   - It's gonna check how many zeros that hash should have at the start
   - 작업증명에서 해쉬 시작에 0이몇개있어야 하는지 확인