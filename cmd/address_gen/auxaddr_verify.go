package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/base58"
)

// Note: If the expected result is a message, then convert the base58 string
//       to binary or hexadecimal and use the desired coding table
//       (ASCII or Unicode) to obtain a plain message.
//
//       In this case we must decode the Skycoin Auxillary address (base58 string)
//       to ASCII and hexadecimal.

func main() {

	addr := os.Args[1]
	aux2hb, _ := base58.Decode(addr)
	/*
		fmt.Println("\033[32m<-----Auxillary Address data----->\033[0m")
		fmt.Println("Decoded base58 hexbytes: ", aux2hb)
		fmt.Println("Decoded base58 address: ", hex.EncodeToString(aux2hb))

		fmt.Println("Bytes Length", len(aux2hb))
		fmt.Println("RipeMD-160 + Version Byte: ", aux2hb[:len(aux2hb)-3])
		fmt.Println("\033[32m<-----Auxillary Address data----->\033[0m")
	*/
	//Status
	fmt.Println("\033[32mRetrieving Skycoin auxillary address parent \033[0m")
	time.Sleep(2 * time.Second)

	hbl := len(aux2hb)
	rmdvb := aux2hb
	regex := regexp.MustCompile(`[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{26,35}`)
	if hbl == 23 {
		hbl = len(aux2hb)
		rmdvb := aux2hb[:len(aux2hb)-2]
		rpvh := cipher.SumSHA256(rmdvb[:])
		chk := rpvh[:4]
		rpvc := append(rmdvb, chk...)
		paddr := base58.Encode(rpvc)
		spaddr := (string(paddr))
		regspaddr := regex.MatchString(spaddr)
		fmt.Println("Parent Skycoin Address: ", string(paddr))
		fmt.Println("Parent address is valid: ", regspaddr)

	} else {
		if hbl == 24 {
			hbl = len(aux2hb)
			rmdvb = aux2hb[:len(aux2hb)-3]
			rpvh := cipher.SumSHA256(rmdvb[:])
			chk := rpvh[:4]
			rpvc := append(rmdvb, chk...)
			paddr := base58.Encode(rpvc)
			spaddr := (string(paddr))
			regspaddr := regex.MatchString(spaddr)
			fmt.Println("Parent Skycoin Address: ", string(paddr))
			fmt.Println("Parent address is valid: ", regspaddr)
		}
	}

}
