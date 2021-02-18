package main

import (
	//"log"
	"fmt"
	"os"
	"bytes"
	"os/exec"
	"strings"
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
	rapip := os.Args[1] //Shellout("ip route show 0.0.0.0/0 dev wlan0 | cut -d\\  -f3")
	path := fmt.Sprintf(`"%s:5454/api/v1/Env"`, rapip)

	//resp := "[{'Error, check ip address of remote AP' err}]"

	_, resp, _ = Shellout(fmt.Sprintf("curl %s", path))

	tl := strings.Trim(resp, "[{")
	tr := strings.Trim(tl, "}]")
	sf := strings.Fields(tr)
	rpk := strings.Trim(sf[0], "visor_pk:")

	fmt.Println(string(rpk))
}
