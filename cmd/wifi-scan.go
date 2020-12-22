package main

import (
    "bytes"
    "fmt"
    "log"
    //"os"
    "flag"
    "os/exec"
    //"strings"
    //"bufio"
)

const ShellToUse = "bash"

func Shellout(command string) (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String(), stderr.String()
}

func main() {
    
    var a string
    flag.StringVar(&a, "i", "wlan0", "wlan interface to scan")
    
    err, out, _ := Shellout(fmt.Sprintf("iwlist %v scan | grep ESSID", a))
    if err != nil {
        log.Printf("error: %v\n", err)
    }

    flag.Parse()
    
    s := fmt.Sprint(out)
    
    fmt.Println(s)
    
}
