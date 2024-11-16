package main

import (
	"github.com/IMHYEWON/hyewoncoin/6.restapi/explorer"
	"github.com/IMHYEWON/hyewoncoin/6.restapi/rest"
)

func main() {
	// goroutine으로 explorer.Start()와 rest.Start() 실행
	go explorer.Start(3000)
	rest.Start(4000)
}
