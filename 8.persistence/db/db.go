package db

import (
	"github.com/IMHYEWON/hyewoncoin/6.restapi/utils"
	"github.com/boltdb/bolt"
)

const dbName = "blockchain.db"
const dataBucket = "data"     // block의 hash를  저장할 bucket
const blocksBucket = "blocks" // block의 정보를 저장할 bucket

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
