package main

import (
	"github.com/IMHYEWON/hyewoncoin/11.wallet/cli"
	"github.com/IMHYEWON/hyewoncoin/11.wallet/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
