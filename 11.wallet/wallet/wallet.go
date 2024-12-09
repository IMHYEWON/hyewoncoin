package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/IMHYEWON/hyewoncoin/11.wallet/utils"
)

// 후에 이 값들은 파일로부터 복원될 것 (임시로 하드코딩)
const privateKey = "93127533446864104697230683079761271464158280997922697724551028140185423594359"
const hashedMessage = "e2d38c3d199d01318e6c0c76693d19a732150acf47e90036c153943e136a4e25"
const signature = "0df690fd9c6bc63f1d870fbc1a503707d8f512d7dd0270e1f5c66ed2e387eb8f98a0e7c568726f842b689b2cc090189b787fd4a7a0a99dfc095d3c68d54e4e37"

func Start() {

	// 1. Private Key 생성 (ECDSA 방식 - )
	// 이후에는 이 Private Key가 파일로부터 복원되어 사용될 것
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	// MarshalECPrivateKey : Private Key를 바이트로 변환
	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)
	fmt.Printf("%x\n", keyAsBytes)
	utils.HandleErr(err)
	fmt.Println("private Key\n", privateKey.D)
	fmt.Println("public Key, x, y\n", privateKey.X, privateKey.Y)

	// 이후에는 이 hash값은 트랜잭션을 hash한 값이 될 것
	hashAsBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	// 2. 서명 생성 (R, S | R: 공개키, S: 개인키)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%x\n", signature)

	utils.HandleErr(err)

}
