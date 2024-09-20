package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/nspcc-dev/neo-go/pkg/vm"
)

const (
	vmHaltedCode    = 0
	wrongArgCode    = 1
	wrongStringCode = 2
	runErrorCode    = 3
	vmFailedCode    = 4
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "expected string (base64) as argument\n")
		os.Exit(wrongArgCode)
	}
	script, err := base64.StdEncoding.DecodeString(string(args[0]))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when decoding string (base64): %s\n", err)
		os.Exit(wrongStringCode)
	}
	vm := vm.New()
	vm.LoadScript(script)
	err = vm.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error when running script: %s\n", err)
		os.Exit(runErrorCode)
	}
	switch {
	case vm.HasFailed():
		fmt.Println("result: VM failed")
		os.Exit(vmFailedCode)
	case vm.HasHalted():
		fmt.Println("result: VM halted")
		fmt.Println(vm.DumpEStack())
		os.Exit(vmHaltedCode)
	}
}
