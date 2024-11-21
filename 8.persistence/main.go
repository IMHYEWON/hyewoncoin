package main

import (
	"github.com/IMHYEWON/hyewoncoin/8.persistence/cli"
	"github.com/IMHYEWON/hyewoncoin/8.persistence/db"
)

func main() {
	// defer : 함수가 끝나기 직전에 실행되는 코드
	defer db.Close()
	cli.Start()
}
