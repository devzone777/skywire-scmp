package main

import (
    "os/exec"
    "flag"
    "log"
    "fmt"
    "os"
)

// Handle flags
var k string

func init() {
    flag.StringVar(&k, "k", "", "\033[33mPubKey to be verified\033[0m")
    flag.String("a", "", "\033[33mGet number of active visors\033[0m")
    
}

func main() {
        flag.Parse()
        pk := string(k)
    	cmd := exec.Command("visorverify", "-k", pk)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}
