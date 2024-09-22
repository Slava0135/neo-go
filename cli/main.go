package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nspcc-dev/neo-go/pkg/vm"
)

type Result struct {
	status string
	errmsg string
	estack string
}

func main() {
	res := run()
	msg := fmt.Sprintf("{\"status\":\"%s\",\"errmsg\":\"%s\",\"estack\":%s}", res.status, res.errmsg, res.estack)
	var out bytes.Buffer
	err := json.Indent(&out, []byte(msg), "", "  ")
	if err != nil {
		panic("error when indenting json: " + err.Error())
	}
	out.WriteTo(os.Stdout)
}

func run() Result {
	args := os.Args[1:]
	if len(args) != 1 {
		return Result{status: "argument error", errmsg: "invalid number of arguments", estack: "[]"}
	}
	script, err := base64.StdEncoding.DecodeString(string(args[0]))
	if err != nil {
		return Result{status: "decoding error", errmsg: fmt.Sprintf("invalid base64 string: %s", err), estack: "[]"}
	}
	vm := vm.New()
	vm.LoadScript(script)
	err = vm.Run()
	if err != nil {
		return Result{status: "VM error", errmsg: fmt.Sprintf("%s", err), estack: "[]"}
	}
	switch {
	case vm.HasHalted():
		return Result{status: "VM halted", errmsg: "", estack: vm.DumpEStack()}
	}
	panic("unknown state")
}
