package main

import (
	"fmt"

	"github.com/pellartech/primelua"
	"github.com/pellartech/primelua/utils"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	// Run examples
	// exampleOne()
	// exampleTwo()
	exampleThree()
}

func exampleOne() {
	// Create a new VM
	vm := primelua.NewVM(1, 1, "", "TSTx8865a8a95769d2479c63df708712df59")

	// Deploy contract
	contractAddr := vm.DeployContract(utils.LoadContractFromFile("./example/contract/example-1.lua"))
	fmt.Println("Contract Address: ", contractAddr)

	// Call Contract concat
	res, _ := vm.CallContract(contractAddr, "concat", 1, lua.LString("start"))

	// Response to get should be "start-concat"
	fmt.Println(res)
}

func exampleTwo() {
	// Create a new VM
	vm := primelua.NewVM(1, 1, "", "TSTx8865a8a95769d2479c63df708712df59")

	// Deploy contract
	contractAddr := vm.DeployContract(utils.LoadContractFromFile("./example/contract/example-2.lua"))
	fmt.Println("Contract Address: ", contractAddr)

	// Call Contract set
	vm.CallContract(contractAddr, "set", 0, lua.LString("test"))

	// Call Contract get
	res, _ := vm.CallContract(contractAddr, "get", 1)

	// Response to get should be "test"
	fmt.Println(res)
}

func exampleThree() {
	// Create a new VM
	vm := primelua.NewVM(1, 1, "e", "TSTx8865a8a95769d2479c63df708712df59")

	// Deploy contract and print address
	contractAddr := vm.DeployContract(utils.LoadContractFromFile("./example/contract/example-3.lua"))
	fmt.Println("Contract Address: ", contractAddr)

	// Mint a new token
	// Minter: TSTx8865a8a95769d2479c63df708712df59
	// Amount: 9999
	// Symbol: TST
	vm.CallContract(contractAddr, "Mint", 0, lua.LString("TSTx8865a8a95769d2479c63df708712df59"), lua.LNumber(9999), lua.LString("TST"))

	// Check token minter address is  "TSTx8865a8a95769d2479c63df708712df59"
	res, _ := vm.CallContract(contractAddr, "getMinter", 1)
	fmt.Println("getMinter: ", res)

	// Check token symbol is "TST"
	res, _ = vm.CallContract(contractAddr, "getSymbol", 1)
	fmt.Println("getSymbol: ", res)

	// Check minter account balance is equal 9999
	res, _ = vm.CallContract(contractAddr, "getAccount", 1, lua.LString("TSTx8865a8a95769d2479c63df708712df59"))
	fmt.Println("getAccount: ", res)

	// Call Tranfer and print response
	// Sender: TSTx8865a8a95769d2479c63df708712df59
	// Recipient: TSTx101
	// Amount: 5
	res, _ = vm.CallContract(contractAddr, "Transfer", 1, lua.LString("TSTx8865a8a95769d2479c63df708712df59"), lua.LString("TSTx101"), lua.LNumber(5))
	fmt.Println("Transfer: ", res)

	// Call getAccount and print response
	res, _ = vm.CallContract(contractAddr, "getAccount", 1, lua.LString("TSTx101"))
	// Response to get should be "5"
	fmt.Println("getAccount: ", res)

	// Call getAccount and print response
	res, _ = vm.CallContract(contractAddr, "getAccount", 1, lua.LString("TSTx8865a8a95769d2479c63df708712df59"))
	// Response to get should be "9994"
	fmt.Println("getAccount: ", res)
}
