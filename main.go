package main

import (
	"flag"
	"fmt"
	"os"
)

var BuildNumber string
var BuildDate string
var GitHash string

func fatal() {
	os.Exit(1)
}

func oops() {
	fmt.Println("usage: opuu <arguments>")
	fmt.Println("'opuu -c help' for help")
	fatal()
}

func main() {
	fmt.Println("opuu! (OnTake Power User Utilities)")
	fmt.Printf("Build Date: %s\n", BuildDate)
	fmt.Printf("Commit: %s\n", GitHash)
	fmt.Printf("Build: %s\n", BuildNumber)
	fmt.Printf("\n")

	command := flag.String("c", "null", "command")
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "directory")
	match := flag.String("m", "null", "match with")
	replace := flag.String("r", "null", "replace with")
	flag.Parse()
	if *command == "http-server" {
		httpserver(port, directory)
	} else if *command == "power-rename" {
		powerrename(directory, match, replace)
	} else {
		oops()
	}
}
