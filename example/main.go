package main

import (
	"fmt"

	"github.com/pellartech/primelua"
	"github.com/pellartech/primelua/utils"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	// Create a new VM
	vm := primelua.NewVM(1, "TSTx8865a8a95769d2479c63df708712df59")

	// Deploy contract
	contractAddr := vm.DeployContract(utils.LoadContractFromFile("./example/contract/example.lua"))
	fmt.Println("Contract Address: ", contractAddr)

	// Call Contract set
	vm.CallContract(contractAddr, "set", 0, lua.LString("test"))

	// Call Contract get
	res, _ := vm.CallContract(contractAddr, "get", 1)

	// Response to get should be "test"
	fmt.Println(res)

	// Call Contract set
	_, _ = vm.CallContract(contractAddr, "set", 0, lua.LString("newvar"))

	// Call Contract get
	res, _ = vm.CallContract(contractAddr, "get", 1)

	// Response to get should be "newvar"
	fmt.Println(res)

	// Call Contract set
	_, _ = vm.CallContract(contractAddr, "concat", 0, lua.LString("start"))

	// Call Contract get
	res, _ = vm.CallContract(contractAddr, "get", 1)

	// Response to get should be "start-concat"
	fmt.Println(res)
}
