package main

import (
	"flag"
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
	fmt.Println(os.Args[2:])

	restCmd := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := restCmd.Int("port", 4000, "Set the port for the server (default: 4000)")

	// os.Args[1] : Command
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Launch the HTML Explorer")
	case "rest":
		restCmd.Parse(os.Args[2:])
	default:
		usage()
	}

	if restCmd.Parsed() {
		fmt.Println("Start the REST API server")
		fmt.Println(*portFlag)
	}
}
