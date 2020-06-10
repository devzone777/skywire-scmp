package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "fmt"
    "log"
)

type activeVisors struct {
    Visor string
    Key string `json:"key"`
    //Start_time  int `json:"start_time"`
}

func main() {
resp, err := http.Get("https://uptime-tracker.skywire.skycoin.com/visors")
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

    s, _ := json.MarshalIndent(c, "", "\t")

    fmt.Println(string(s))
}
