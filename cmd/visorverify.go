package main

import (
    //"github.com/go-validator/validator"
    //"gopkg.in/validator.v2"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "strings"
    //"errors"

    //"reflect"
    //"unsafe"
    //"bytes"

    "flag"
    "fmt"
    "log"
)

// Filter for specific keys in visor data
type activeVisors struct {
    Visor string
    Key string `json:"key"`
    //Start_time  int `json:"start_time"`
}

// Handle flag and point it's pointer to it's string
var pk string

func init() {
    flag.StringVar(&pk, "k", "", "\033[33mPubKey to be verified\033[0m")
}

func main() {
    // Retrieve activevisor data
    resp, err := http.Get("https://uptime-tracker.skywire.skycoin.com/visors")
    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    // Move resp.Body into body 
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    // Convert http resp to raw json
    var jsonkeyvalues []activeVisors
    err = json.Unmarshal(body, &jsonkeyvalues)
    if err != nil {
        log.Fatal(err)
    }

    // Format to proper json structure
    s, _ := json.MarshalIndent(jsonkeyvalues, "", "\t")

    // Stringify json data
    str := string(s)

    // Set flags
    flag.Parse()

    // Verify valid pk and convert struct to string
    //var vpk []visorKey
    //err = json.Unmarshal(pk, &vpk)
    //if err != nil {
        //log.Fatal(err)
    //}
    
    //cvpk := string(vpk)

    //fmt.Println("len(s)", len(s))
    l := len(pk)
    if l != 66 {
            log.Fatal(err)
        }

    // Search for visor and show result
    if strings.Contains(str, pk) {
        fmt.Println("\033[32mVisor is active")
        fmt.Println("\033[0m")
        } else {
        fmt.Println("\033[32mVisor is NOT active")
        fmt.Println("\033[0m")
    }
    fmt.Println("\033[0m")
}
