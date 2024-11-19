package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Welcome to hyewon's CLI")
	fmt.Println("Please use the following commands:")
	fmt.Println("explorer: Launch the HTML Explorer")
	fmt.Println("rest: Launch the REST API (recommended)")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	// os.Args[1] : Command
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Launch the HTML Explorer")
	case "rest":
		fmt.Println("Launch the REST API")
	default:
		usage()
	}
}
