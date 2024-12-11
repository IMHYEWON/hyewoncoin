package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/IMHYEWON/hyewoncoin/10.transaction/utils"
)

const walletFileName string = "hyewoncoin.wallet"

type wallet struct {
	privateKey *ecdsa.PrivateKey
	address    string // it will be a PublicKey
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
		w.address = addressFromPrivateKey(w.privateKey)
	}
	return w
}
