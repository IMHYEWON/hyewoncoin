package blockchain

import "sync"

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockchain struct {
	blocks []block
}

// blockchain.go 내부에서만 이 변수에 접근 가능
var b *blockchain

// sync.Once : 한번만 실행되는 코드를 실행하기 위한 구조체
var once sync.Once

// 한번만 실행되는 코드
// 이 메소드로 블록체인 생성을 제어
func GetBlockChain() *blockchain {
	// nil : 아무것도 없음을 나타내는 특별한 값
	// b가 nil이면 새로운 blockchain을 생성하여 반환
	if b == nil {
		// 블록체인 인스턴스 생성이 단 한번만 실행되도록 보장
		once.Do(func() {
			b = &blockchain{}
		})
	}

	// 블록체인 인스턴스 반인
	return b
}
