package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/nspcc-dev/neo-go/pkg/vm"
	"github.com/nspcc-dev/neo-go/pkg/vm/opcode"
)

type Result struct {
	status string
	errmsg string
	lastop byte
	estack string
}

func main() {
	res := run()
	msg := fmt.Sprintf("{\"status\":\"%s\",\"errmsg\":\"%s\",\"lastop\":%d,\"estack\":%s}", res.status, res.errmsg, res.lastop, res.estack)
	fmt.Println(msg)
	var out bytes.Buffer
	err := json.Indent(&out, []byte(msg), "", "  ")
	if err != nil {
		panic("error when indenting json: " + err.Error())
	}
	out.WriteTo(os.Stdout)
}

func run() Result {
	args := os.Args[1:]
	var lastop byte = 0
	if len(args) != 1 {
		return Result{status: "argument error", errmsg: "invalid number of arguments", lastop: lastop, estack: "[]"}
	}
	script, err := base64.StdEncoding.DecodeString(string(args[0]))
	if err != nil {
		return Result{status: "decoding error", errmsg: fmt.Sprintf("invalid base64 string: %s", err), lastop: lastop, estack: "[]"}
	}
	vm := vm.New()
	vm.SetOnExecHook(func(scriptHash util.Uint160, offset int, opcode opcode.Opcode) {
		lastop = byte(opcode)
	})
	vm.LoadScript(script)
	err = vm.Run()
	if err != nil {
		return Result{status: "VM error", errmsg: fmt.Sprintf("%s", err), lastop: lastop, estack: "[]"}
	}
	switch {
	case vm.HasHalted():
		return Result{status: "VM halted", errmsg: "", lastop: lastop, estack: vm.DumpEStack()}
	}
	panic("unknown state")
}
