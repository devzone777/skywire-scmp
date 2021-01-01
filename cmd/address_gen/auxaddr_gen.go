package main

import (
	"time"
        "fmt"
        "os"
        "github.com/skycoin/skycoin/src/cipher/base58"

)

// Note: If the expected result is a message, then convert the base58 string 
//       to binary or hexadecimal and use the desired coding table 
//       (ASCII or Unicode) to obtain a plain message. 
//
//       In this case we must decode the Skycoin address (base58 string)
//       to hexadecimal. 




func main() {
	
	addr := os.Args[1]
	dba, _ := base58.Decode(addr)
	conv1 := dba[:len(dba)-2]
	fmt.Println("\033[32mGenerating new Auxillary Address\033[0m")
	time.Sleep(2 * time.Second)
	fmt.Println("Parent Skycoin Address: ", addr)
	fmt.Println("New Auxillary address : ", base58.Encode(conv1))

}

