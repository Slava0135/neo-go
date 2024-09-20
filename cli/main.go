package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "expected string (base64) as argument\n")
		os.Exit(1)
	}
	script, err := base64.StdEncoding.DecodeString(string(args[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when decoding string (base64): %s\n", err)
		os.Exit(2)
	}
	fmt.Printf("%v\n", script)
}
