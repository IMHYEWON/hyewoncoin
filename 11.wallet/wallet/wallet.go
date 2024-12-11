package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/IMHYEWON/hyewoncoin/10.transaction/utils"
)

const walletFileName string = "hyewoncoin.wallet"

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string // it will be a PublicKey
}

var w *wallet

func hasWalletFile() bool {
	// 파일이 있는지 확일
	_, err := os.Stat(walletFileName)

	// 파일이 없으면 false, 있으면 true
	return os.IsExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	// private key 생성
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	fmt.Println("Private Key: ", privateKey)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	// x509 : encoding/asn1, encoding/pem, crypto/x509
	// MarshalECPrivateKey : private key를 byte로 변환
	// MarshalECPrivateKey converts an EC private key to SEC 1, ASN.1 DER form.
	// This kind of key is commonly encoded in PEM blocks of type "EC PRIVATE KEY". For a more flexible key format which is not EC specific, use [MarshalPKCS8PrivateKey].
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	fmt.Println("Key Bytes: ", bytes)

	err = os.WriteFile(walletFileName, bytes, 0644)
	utils.HandleErr(err)
}

func restoreKey() *ecdsa.PrivateKey {
	keyAsBytes, err := os.ReadFile(walletFileName)
	utils.HandleErr(err)

	key, err := x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)
	return key
}

// private Key로 부터 address를 생성
func addressFromPrivateKey(key *ecdsa.PrivateKey) string {
	z := append(key.X.Bytes(), key.Y.Bytes()...)
	return fmt.Sprintf("%x", z)
}

func restoreBigInts(signature string) (*big.Int, *big.Int, error) {
	// 1. signature string을 byte로 변환
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		return nil, nil, err
	}

	// 2. signature를 반으로 나눠서 r, s로 변반
	firstHalfBytes := signatureBytes[:len(signatureBytes)/2]
	seconHalfBytes := signatureBytes[len(signatureBytes)/2:]

	// 3. r, s를 big.Int로 변환
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(seconHalfBytes)

	return &bigA, &bigB, nil
}

// sign : payload를 받아서 private key로 sign한 결과를 리턴
// payload : transction 데이터
// When we sign a transaction, anything has not been encrypted (we dont change the data)
// We just create a signature for the data
// It means that we can verify the data with the signature
// So When we verify the data, we should get the 'data' and the 'signature' both
func sign(payload string, w *wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleErr(err)

	signature := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", signature)
}

// verify : signature, payload, address를 받아서 검증
// we dont need the private key to verify the data
func verify(signature string, payload string, address string) bool {
	// 1. restore signature
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)

	// 2. restore public key
	x, y, err := restoreBigInts(address)
	utils.HandleErr(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}

	// 3. payload를 byte로 변환
	payloadBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	// 4. verify
	ok := ecdsa.Verify(&publicKey, payloadBytes, r, s)
	return ok
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		// if user ahs a wallet already?
		// yes -> restore from file(db)
		// no -> create private key, save to file(db)
		if hasWalletFile() {
			// restore
			w.privateKey = restoreKey()
		} else {
			// create private key
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = addressFromPrivateKey(w.privateKey)
	}
	return w
}
