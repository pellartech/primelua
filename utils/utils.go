package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

// GenerateContractHash - Returns the hash/address for
// the contract e.g 8865a8a95769d2479c63df708712df59
func GenerateContractHash(contract string) string {
	hash := md5.Sum([]byte(contract))
	return hex.EncodeToString(hash[:])
}

// LoadContractFromFile - Read a contract from file
// and return it as a string
func LoadContractFromFile(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
