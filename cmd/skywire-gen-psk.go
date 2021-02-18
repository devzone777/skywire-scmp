package main

import (
	"os"
	"fmt"
	"github.com/skycoin/skycoin/src/cipher/base58"
)

func main() {
	auxaddr := os.Args[1]
	psk := base58.Encode([]byte(auxaddr))
	fmt.Println(psk)
}
