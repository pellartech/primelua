package primelua

import (
	"errors"
	"fmt"

	"github.com/pellartech/primelua/database"
	"github.com/pellartech/primelua/utils"
	lua "github.com/yuin/gopher-lua"
)

// Vars - Stores the vars from the contracts
var Vars map[string]string

// VM - Stores a lua vm state
type VM struct {
	State        *lua.LState
	BlockIndex   int
	OwnerAddress string
}

// NewVM - Creates a new VM
func NewVM(blockIndex int, ownerAddress string) VM {
	vm := VM{}
	vm.State = lua.NewState()
	vm.BlockIndex = blockIndex
	vm.OwnerAddress = ownerAddress
	Vars = make(map[string]string)
	vm.State.SetGlobal("get_value", vm.State.NewFunction(getValue))
	vm.State.SetGlobal("set_value", vm.State.NewFunction(setValue))
	return vm
}

// DeployContract - deploys a new contract to
// the blockchain. Returns the contract address
func (luavm *VM) DeployContract(contract string) string {

	// Generate the address for the contract
	contractAddress := utils.GenerateContractHash(contract)

	// Store the contract in ContractDB
	err := database.ContractDB.Put([]byte(contractAddress), []byte(contract), nil)
	if err != nil {
		fmt.Println("DeployContract1 - Error: Writing contract to db")
	}

	return contractAddress
}

// CallContract - call a new contract from
// the blockchain.
func (luavm *VM) CallContract(contractAddr string, function string, ret int, inputs ...lua.LValue) (lua.LString, error) {

	// Get the contract from the db
	contractCode, err := database.ContractDB.Get([]byte(contractAddr), nil)
	if err != nil {
		return "", err
	}

	// Read contract from db
	err = luavm.State.DoString(string(contractCode))
	if err != nil {
		return "", err
	}

	if err := luavm.State.CallByParam(lua.P{
		Fn:      luavm.State.GetGlobal(function), // name of Lua function
		NRet:    ret,                             // number of returned values
		Protect: true,                            // return err or panic
	}, inputs...); err != nil {
		return "", err
	}

	// Get the returned value from the stack and cast it to a lua.LString
	if str, ok := luavm.State.Get(-1).(lua.LString); ok {
		// Pop the returned value from the stack
		luavm.State.Pop(1)
		// return the value from top of stack
		return str, nil
	}
	return "", errors.New("end of CallContract reached")
}

// setValue - Sets the local state vars in go
// from the lua contract
func setValue(L *lua.LState) int {
	name := L.ToString(1)
	value := L.ToString(2)
	Vars[name] = value
	return 0
}

// getValue - Gets the local state vars in go
// and sets them in lua contract
func getValue(L *lua.LState) int {
	name := L.ToString(1)
	L.Push(lua.LString(Vars[name]))
	return 1
}
