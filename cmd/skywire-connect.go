package main

import (
	"os"
	"fmt"
	"bytes"
	"log"
	"os/exec"
	"strings"
	
	"github.com/skycoin/skycoin/src/cipher/base58"
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
	
	a := os.Args[1]
	fmt.Println(a)
	
	fmt.Println("debug:ssid=", a)
	fmt.Println("debug:SSID Length=", len(a))
	ssids := a
	fmt.Println("debug:ssids=", ssids)
	
	_, network, _ := Shellout("wpa_cli add_network") //#1
	snetwork := strings.Split(network, "\n")
	fmt.Println("debug:network=", snetwork[1])
	netnum := snetwork[1]
	fmt.Println("debug:netnum=", netnum)
	psk := base58.Encode([]byte(ssids))
	fmt.Println("debug:psk=", psk)
	fmt.Println("debug:PSK Bytes Length=", len(psk))

	//Set SSID #2
	sn := fmt.Sprintf("wpa_cli set_network %s ssid '\"%s\"'", netnum, ssids)
	fmt.Println("debug:sn=", sn)
	setnet := string(sn)
	fmt.Println("debug:setnet=", setnet)
	Shellout(setnet)

	//Set PSK #4
	pk := fmt.Sprintf("wpa_cli set_network %s psk '\"%s\"'", netnum, psk)
        fmt.Println("debug:pk=", pk)
        setpsk := string(pk)
        fmt.Println("debug:setpsk=", setpsk)
	Shellout(setpsk)

	//Set security type #5
	sec := fmt.Sprintf("wpa_cli set_network %s key_mgmt WPA-PSK", netnum)
        fmt.Println("debug:security type=", sec)
        setsecurity := string(sec)
        fmt.Println("debug:setnet=", setsecurity)
        Shellout(setsecurity)


	//Set priority #6
        pr := fmt.Sprintf("wpa_cli set_network %s priority %s", netnum, netnum)
        fmt.Println("debug:priority=", pr)
        setpr := string(pr)
        fmt.Println("debug:setpr=", setpr)
        Shellout(setpr)
	

	//Enable network #7
        en := fmt.Sprintf("wpa_cli enable_network %s", netnum)
        fmt.Println("debug:enable=", en)
        enable := string(en)
        fmt.Println("debug:setsel=", enable)
        Shellout(enable)

	//Save settings #8
	err, status, _ := Shellout("wpa_cli save")
	fmt.Println(status)
	if err != nil {
            log.Printf("error: %v\n", err)

	//Select network #9
        sn := fmt.Sprintf("wpa_cli select_network %s", netnum)
        fmt.Println("debug:select=", sn)
        selnet := string(sn)
        fmt.Println("debug:selnet=", selnet)
        Shellout(selnet)
	/*
	//Reboot #10
        rb := fmt.Sprintf("reboot")
        reboot := string(rb)
        Shellout(reboot)
	*/

	Shellout("pkill wpa_supplicant & screen -dmS skywire-wpa-supplicant wpa_supplicant -c /etc/wpa_supplicant/wpa_supplicant.conf -i wlan0")

	}

//exec.Command("reboot")
}
