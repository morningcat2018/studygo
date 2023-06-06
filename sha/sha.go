package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

func main() {

	bytes := []byte("世界人民大团结万岁")

	h := sha256.New()
	h.Write(bytes)
	fmt.Printf("%x\n", h.Sum(nil))

	var h512 hash.Hash = sha512.New()
	h512.Write(bytes)
	fmt.Printf("%x\n", h512.Sum(nil))

}
