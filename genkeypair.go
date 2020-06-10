// Package cipher implements common golang encoding interfaces for
// github.com/SkycoinProject/skycoin/src/cipher
package main

import (
        "fmt"

        "github.com/SkycoinProject/dmsg/cipher"
)

func main() {
        //Print new Private Key and Secret Key
        pk, sk := cipher.GenerateKeyPair()
        fmt.Println("New Pubkey:", pk)
        fmt.Println("New Seckey:", sk)
}

