package utils

import (
	"bytes"
	"encoding/gob"
)

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// interface{} : 모든 타입을 나타내는 빈 인터페이스
// Java의 Object와 비슷한 역할
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))

	return aBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}
