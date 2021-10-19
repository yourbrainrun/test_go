package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := sha256.New()
	hash.Write([]byte("tests"))

	myBytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(myBytes)
	fmt.Println(hashCode)

	bytesSli := sha256.Sum256([]byte("tests"))
	hashCode1 := hex.EncodeToString(bytesSli[:])

	fmt.Println(hashCode1)
	fmt.Println(fmt.Sprintf("%x", bytesSli))
}
