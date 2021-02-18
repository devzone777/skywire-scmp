/*
swap app for skywire visor
*/
package main

import (
	"flag"
	"fmt"
	"os/exec"
	"net/http"
	"strings"
	"bytes"
	"log"
	"github.com/gorilla/mux"
	
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


var addr = flag.String("addr", ":5454", "address to bind")

func main() {
	
	flag.Parse()
	log.Println("Successfully started swap service")

	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	//api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("/Env", env)

	log.Println("Serving SWAP ENV Params on localhost:5454")
	log.Fatal(http.ListenAndServe(":5454", r))
}

func env(w http.ResponseWriter, r *http.Request) {
    
    _, mac, _ := Shellout("ifconfig wlan1 | grep ether")
    macs := fmt.Sprint(mac)
    f := strings.Fields(macs)
    macf := f[1]
    tl := strings.Trim(macf, "dc:")
    
    _, VISOR_PK, _ := Shellout("skywire-cli visor pk")
    _, AP_ADDR, _ := Shellout("ip route show 0.0.0.0/0 dev wlan1 | cut -d\\  -f3")
    MAC_ADDR := tl

    vv := fmt.Sprintf("{visor_pk:%s", VISOR_PK)
    pp := fmt.Sprintf("ap_addr:%s", AP_ADDR)
    kk := fmt.Sprintf("mac_addr:%s}", MAC_ADDR)
    rstr := fmt.Sprintln(vv, pp, kk)
    rstrs := strings.Fields(rstr)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintln(w, rstrs)
}

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get not used yet"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post not used yet"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put not used yet"}`))
}

//func delete(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json")
    //w.WriteHeader(http.StatusOK)
    //w.Write([]byte(`{"message": "delete not used yet"}`))
//}

