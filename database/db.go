package database

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// ContractDB - Instnce of DB
var ContractDB *leveldb.DB

// StateDB - Instnce of DB
var StateDB *leveldb.DB

func init() {
	// The returned DB instance is safe for concurrent use. Which mean that all
	// DB's methods may be called concurrently from multiple goroutine.
	var err error
	ContractDB, err = leveldb.OpenFile("data/contracts", nil)
	if err != nil {
		fmt.Println("Failed to open contracts database")
		panic(err)
	}
	StateDB, err = leveldb.OpenFile("data/states", nil)
	if err != nil {
		fmt.Println("Failed to open states database")
		panic(err)
	}
}
