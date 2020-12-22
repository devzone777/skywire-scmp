package main

import (
    "bytes"
    "fmt"
    "log"
    //"os"
    "flag"
    "os/exec"
    "strings"
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
    ss := string(s)
    sf := strings.Fields(ss)
    fmt.Println("ssid 1 length: ", len(sf[1]))
    if len(sf[1])==40 {
        fmt.Println("\033[36m1. \033[32m", sf[1])
    }
    fmt.Println("\033[0mssid 2 length: ", len(sf[2]))
    if len(sf[2])==40 {
        fmt.Println("\033[36m2. \033[32m", sf[2])
    }
    fmt.Println("\033[0mssid 3 length: ", len(sf[3]))
    if len(sf[3])==40 {
        fmt.Println("\033[36m3. \033[32m", sf[3])
    }
    fmt.Println("\033[0mssid 4 length: ", len(sf[4]))
    if len(sf[4])==40 {
        fmt.Println("\033[36m4. \033[32m", sf[4])
    }
    fmt.Println("\033[0mssid 5 length: ", len(sf[5]))
    if len(sf[5])==40 {
        fmt.Println("\033[36m5. \033[32m", sf[5])
        fmt.Println("\033[0m")
    }
    
}
