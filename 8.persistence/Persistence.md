# 8. Persistence
- Key, Value DB인 BOLT를 사용
  - Key : hash value, Value : Actual Data
  - Bolt DB 의 API는 간단하고, Stable하기 때문에 바뀔 가능성이 적음

# 8.1 Creating the Database
- Bolt DB 설치
  -`go get github.com/boltdb/bolt`
- DB는 싱글톤으로 관리하며, 없을 경우 생성
  - block의 k,v data를 저장할 bucket과 블록의 메타정보를 저장할 버킷 생서
  
# 8.2 A New Blockchain
- 메모리에 Block Slice들을 모두 저장할 필요가 없기 때문에 `blockchain.go`의 수정이 필요함
  - 이제부터는 blockchain에 block을 추가할 때 slice를 append할 필요 없음
  - 블록 생성은 이전과 동일 (함수 분리)
  - 블록을 생성하고 나서는, 생성된 블록(=블록 체인의 마지막 블록)의 hash와 블록체인의 height만 저장
    - DB에 블록들을 저장할 예정

# 8.3 Saving A Block
- 블록을 실제DB에 저장
  - db.go 는 Data access Layer (Interface)의 역할을 할 예정  

# 8.4 Persisting the Blockchain
- Blockchain에 Block이 Add될 때 마지막 Block을 체크포인트로 DB에 저장
  - 이후 이 체크포인트로 마지막 블록을 불러와 확인