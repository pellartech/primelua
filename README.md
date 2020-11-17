# primelua
Smart contracts for pellar-prime

### Example Smart contract being called below
```lua
function concat(a)
	a = a .. "-concat"
	set_value("a", a)
	return get_value("a")
end
```

### Example
```go
// Create a new VM
vm := primelua.NewVM(1, "TSTx8865a8a95769d2479c63df708712df59")

// Deploy contract
contractAddr := vm.DeployContract(utils.LoadContractFromFile("./example/contract/example.lua"))
fmt.Println("Contract Address: ", contractAddr)

// Call Contract set
_, _ = vm.CallContract(contractAddr, "concat", 0, lua.LString("start"))

// Call Contract get
res, _ = vm.CallContract(contractAddr, "get", 1)

// Response to get should be "start-concat"
fmt.Println(res)
 ```
 
### Output
```bash
Contract Address:  c462033b1f2ffe2681d61f9d9e2a63bd

start-concat
```
