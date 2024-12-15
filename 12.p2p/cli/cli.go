package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/IMHYEWON/hyewoncoin/12.p2p/explorer"
	"github.com/IMHYEWON/hyewoncoin/12.p2p/rest"
)

func usage() {
	fmt.Println("Welcome to hyewon's CLI")
	fmt.Println("Please use the following commands:")
	fmt.Println("-port=4000: Launch the HTML Explorer")
	fmt.Println("-mode=rest: Launch the REST API (recommended)")
	runtime.Goexit() // it will termainate all program, but it will not terminate the defer function
}
func Start() {
	if len(os.Args) < 2 {
		usage()
	}

	port := flag.Int("port", 3000, "port to run the server on")
	mode := flag.String("mode", "rest", "Choose between 'rest' and 'explorer' (default 'rest')")

	// os.Args[1:]로 시작하는 플래그 값을 설정
	flag.Parse()

	switch *mode {
	case "rest":
		fmt.Println("Starting REST API on port", *port)
		rest.Start(*port)
	case "explorer":
		fmt.Println("Starting HTML Explorer ")
		explorer.Start()
	case "both":
		fmt.Println("Starting REST API and HTML Explorer on port", *port)
		go rest.Start(*port)
		explorer.Start()
	default:
		usage()
	}

	fmt.Println(*port, *mode)
}
