package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "fmt"
    "log"
)

type activeVisors struct {
    Key string `json:"key"`
    Start_time  int `json:"start_time"`
}

func main() {
resp, err := http.Get("https://uptime-tracker.skywire.skycoin.com/uptimes")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    var c []activeVisors
    err = json.Unmarshal(body, &c)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("\033[32mActive visor count =", len(c))
    fmt.Println("\033[0m")
}
