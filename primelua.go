package primelua

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/pellartech/primelua/database"
	"github.com/pellartech/primelua/utils"
	lua "github.com/yuin/gopher-lua"
)

// VM - Stores a lua vm state
type VM struct {
	State           *lua.LState
	BlockIndex      int
	BlockTimestamp  int
	ContractAddress string
	SenderAddress   string
	// TempStates - Stores the vars from the contracts
	TempStates map[string]string
}

// NewVM - Creates a new VM
func NewVM(blockIndex int, blockTimestamp int, contractAddress string, senderAddress string) VM {
	// Build object
	vm := VM{}
	vm.State = lua.NewState()
	vm.BlockIndex = blockIndex
	vm.BlockTimestamp = blockTimestamp
	vm.ContractAddress = contractAddress
	vm.SenderAddress = senderAddress
	// Init temp state storage
	vm.TempStates = make(map[string]string)
	// Set build in functions
	vm.setGlobals()
	return vm
}

// setGlobals - Add built in functions to VM
func (vm *VM) setGlobals() {
	vm.State.SetGlobal("get_value", vm.State.NewFunction(vm.getValue))
	vm.State.SetGlobal("set_value", vm.State.NewFunction(vm.setValue))
	vm.State.SetGlobal("block_index", lua.LNumber(vm.BlockIndex))
	vm.State.SetGlobal("sender_address", lua.LString(vm.SenderAddress))
	vm.State.SetGlobal("contract_address", lua.LString(vm.ContractAddress))
}

// DeployContract - deploys a new contract to
// the blockchain. Returns the contract address
func (vm *VM) DeployContract(contract string) string {
	// Generate the address for the contract
	contractAddress := utils.NewRandomAddress()

	// Store the contract in ContractDB
	err := database.ContractDB.Put([]byte(contractAddress), []byte(contract), nil)
	if err != nil {
		fmt.Println("DeployContract1 - Error: Writing contract to db")
	}

	return contractAddress
}

// CallContract - call a new contract from
// the blockchain.
func (vm *VM) CallContract(contractAddr string, function string, ret int, inputs ...lua.LValue) (lua.LString, error) {
	// Get the contract from the db
	contractCode, err := database.ContractDB.Get([]byte(contractAddr), nil)
	if err != nil {
		return "", err
	}

	// Get the contract states from the db
	states, err := database.StateDB.Get([]byte(contractAddr), nil)
	if err == nil {
		buf := bytes.NewBuffer(states)
		d := gob.NewDecoder(buf)
		err = d.Decode(&vm.TempStates)
		if err != nil {
			panic(err)
		}
	}

	// Read contract from db
	err = vm.State.DoString(string(contractCode))
	if err != nil {
		return "", err
	}

	if err := vm.State.CallByParam(lua.P{
		Fn:      vm.State.GetGlobal(function), // name of Lua function
		NRet:    ret,                          // number of returned values
		Protect: true,                         // return err or panic
	}, inputs...); err != nil {
		return "", err
	}

	// Get the returned value from the stack and cast it to a lua.LString
	if str, ok := vm.State.Get(-1).(lua.LString); ok {
		// Update states
		buf := bytes.NewBuffer(nil)
		enc := gob.NewEncoder(buf)
		err := enc.Encode(&vm.TempStates)

		// Get the contract states from the db
		err = database.StateDB.Put([]byte(contractAddr), buf.Bytes(), nil)
		if err != nil {
			return "", errors.New("Error updating contract  states")
		}
		// Pop the returned value from the stack
		vm.State.Pop(1)
		// return the value from top of stack
		return str, nil
	}
	return "", errors.New("end of CallContract reached")
}

// setValue - Sets the local state vars in go
// from the lua contract
func (vm *VM) setValue(L *lua.LState) int {
	name := L.ToString(1)
	value := L.ToString(2)
	vm.TempStates[name] = value
	return 0
}

// getValue - Gets the local state vars in go
// and sets them in lua contract
func (vm *VM) getValue(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LString(vm.TempStates[name]))
	return 1
}
