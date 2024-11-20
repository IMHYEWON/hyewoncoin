# 8. Persistence
- Key, Value DB인 BOLT를 사용
  - Key : hash value, Value : Actual Data
  - Bolt DB 의 API는 간단하고, Stable하기 때문에 바뀔 가능성이 적음

# 7.1 Creating the Database
- Bolt DB 설치
  -`go get github.com/boltdb/bolt`
- DB는 싱글톤으로 관리하며, 없을 경우 생성
  - block의 k,v data를 저장할 bucket과 블록의 메타정보를 저장할 버킷 생서