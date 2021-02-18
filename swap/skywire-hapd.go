package main

import (
	"os"
	"fmt"
	"bytes"
	"os/exec"
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
	config := fmt.Sprintf("country_code=US\ninterface=wlan1\n#If this fails, try rt1871xdrv \ndriver=nl80211\n# Name of host skywire network\nssid=%s\n# Pick a channel not already in use\nchannel=6\n# Change to b for older devices\nhw_mode=g\nmacaddr_acl=0\nauth_algs=3\n# Disable this to insure the AP is visible:\nignore_broadcast_ssid=0\nwpa=2\nwpa_passphrase=\"%s\"\nwpa_key_mgmt=WPA-PSK\nwpa_pairwise=TKIP\nrsn_pairwise=CCMP", os.Args[1], os.Args[2])
	fmt.Println(config)
}
