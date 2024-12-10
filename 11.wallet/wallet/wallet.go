package wallet

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/IMHYEWON/hyewoncoin/11.wallet/utils"
)

// 후에 이 값들은 파일로부터 복원될 것 (임시로 하드코딩)
const signature string = "a30f035c8d44dea5c1355a963e18e0a9426622cb3467f8b553538099ffab6fabc775b6d9a9014cd01be0436c980cf6deec7e4cacaa0258d7ecc3e77c329ab773"
const privateKey string = "307702010104204eb9438b09d338be90fdc55d9391d4d5e10f620ca458818a3941e0bd400c2514a00a06082a8648ce3d030107a14403420004cf39c60bcf171dfdb10d9d51345558c522169e23f4c148c80658cac50303f0e7d410fe4990f2835793ab6fc5f53d81260346deb6dc38d2041f35d88c31f7fd9c"
const hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"

func Start() {

	// 1. PrivateKey 복원
	// 1-A. pricateKey가 hex 포맷이 맞는지 먼저 체크 (아무도 이 키를 조작하지 않았는지 확인하기 위해 인코딩해서 확인)
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	// 1-B. 비공개키를 x509.ParseECPrivateKey() 함수로 복원
	restoredKey, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	fmt.Println(restoredKey)

	// 2. Transaction 검증 (서명 검증)
	// signature :
	// 	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	// 	signature := append(r.Bytes(), s.Bytes()...)
	// 	위 코드에서 r, s를 append한 값이 signature
	// 	따라서 signature를 두 부분으로 나눠서 r, s로 복원
	sigBytes, err := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]
	fmt.Printf("%d\n\n%d\n\n%d\n\n", sigBytes, rBytes, sBytes)

	// 2-A. r, s를 big.Int로 변환 (서명 복원)
	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
	/* Output:
	-- false : 양수
	{false [6004284127122780075 4784549910280665269 13922133424419168425 11749613648874888869]}
	{false [17060734334020532083 17041142344486901975 2008679567493756638 14372594831782399184]}
	*/

	// 2-B. 서명된 해시메시지를 DecodeString() 함수로 바이트로 변환
	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	// 3. 공개키를 이용해 검증 (서명이 유효한지 확인)
	/*
		Verify verifies the signature in r, s of hash using the public key, pub.
		Its return value records whether the signature is valid.
		Most applications should use VerifyASN1 instead of dealing directly with r, s.
	*/
	ok := ecdsa.Verify(&restoredKey.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok) // true

}
