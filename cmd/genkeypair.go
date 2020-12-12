// Package cipher implements common golang encoding interfaces for
// github.com/Skycoin/skycoin/src/cipher
package main

import (
        "fmt"

        "github.com/Skycoin/dmsg/cipher"
)

func main() {
        //Print new Private Key and Secret Key
        pk, sk := cipher.GenerateKeyPair()
        fmt.Println("\033[32mNew Pubkey:\033[0m", pk)
        fmt.Println("\033[032mNew Seckey:\033[0m", sk)
}

