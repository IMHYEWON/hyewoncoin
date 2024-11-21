package db

import (
	"fmt"

	"github.com/IMHYEWON/hyewoncoin/8.persistence/utils"
	"github.com/boltdb/bolt"
)

const dbName = "blockchain.db"
const dataBucket = "data"     // block의 hash를  저장할 bucket
const blocksBucket = "blocks" // block의 정보를 저장할 bucket
const checkpoint = "checkpoint"

var db *bolt.DB

// Singleton Pattern을 활용해서 db를 한번만 생성하고 사용, 실제 db변수는 외부에서 접근할 수 없도록 private으로 선언
// initialize the database
func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		utils.HandleErr(err)
		db = dbPointer

		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)

			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

// Block Bucket에 Block을 저장, key: hash, value: block data
// data는 []byte로 받아서 저장 (block.go에서 toBytes()로 []byte로 변환)
func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)

	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data) // key: hash, value: block data
		return err
	})

	utils.HandleErr(err)
}

// 블록체인에는 마지막 블록의 hash만 저장
func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var data []byte

	fmt.Println("Getting Checkpoint...")
	// DB에서 checkpoint를 가져와서 반환
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})

	fmt.Printf("Data: %x\n", data)
	return data
}
