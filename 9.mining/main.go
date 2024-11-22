package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

func main() {
	difficulty := 6
	target := strings.Repeat("0", difficulty)
	nonce := 1 // nonce란 ? : 블록체인에서 사용되는 숫자로, 블록을 찾기 위해 계속 변경되는 숫자

	start := time.Now()

	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		// fmt.Printf("%x\n", hash)

		// hash값이 target값으로 시작하는지 확인
		if strings.HasPrefix(hash, target) {
			fmt.Printf("Found hash: %s (Target: %s)\n nonce: %d\n\n", hash, target, nonce)
			break
		} else {
			nonce++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed Time: %s\n", elapsed)
}
