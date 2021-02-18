package main

import (
	"fmt"
	"os"
	"time"

	"github.com/skycoin/skycoin/src/cipher/base58"
)

// Note: If the expected result is a message, then convert the base58 string
//       to binary or hexadecimal and use the desired coding table
//       (ASCII or Unicode) to obtain a plain message.
//
//       In this case we must decode the Skycoin address (base58 string)
//       to hexadecimal.

func main() {

	//Take address from cmd argument
	addr := os.Args[1]

	//Handle debug args
	arg2 := ""
	if os.Args != nil && len(os.Args) > 2 {
		arg2 = os.Args[2]
	}

	//Decode base 58 address to hexadecimal bytes
	dba, _ := base58.Decode(addr)

	//Check length of Skycoin address
	salen := fmt.Sprint(len(addr))
	if arg2 == "debug" {
		fmt.Println("\033[32m<------Input address info------>\033[0m\nInput address length: ", salen)
	}
	conv1 := dba

	//Handle address length
	if salen == "35" {
		conv1 = dba[:len(dba)-2]
	} else {
		if salen == "34" {
			conv1 = dba[:len(dba)-2]
		} else {
			if salen == "33" {
				//dba[24] = dba[23]
				//conv1 = dba[:len(dba)]

				conv1 = dba[:len(dba)-1]
			}
		}
	}

	//Debug output
	if arg2 == "debug" {
		fmt.Println("Input address hexbytes: ", dba)
	}

	//Status
	//fmt.Println("\033[32mGenerating new Auxillary Address\033[0m")
	time.Sleep(2 * time.Second)

	fmt.Println("Parent Skycoin Address: ", addr)

	if arg2 == "debug" {
		fmt.Println("Parent Skycoin Address Length", len(addr))
	} else {
	}

	fmt.Println("New Auxillary address: ", base58.Encode(conv1))

	//Debug output
	if arg2 == "debug" {
		fmt.Println("Auxillary Address Length", len(base58.Encode(conv1)))
	} else {
	}
}
