package main

import (
	"bytes"
	"fmt"
	"log"
	//"os"
	"flag"
	"time"
	"os/exec"
	"strings"
	"regexp"

	//"github.com/skycoin/skycoin/src/cipher/base58"
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

	var arg2 string
	flag.StringVar(&arg2, "m", "", "Select mode [debug | verbose]")

	err, out, _ := Shellout(fmt.Sprintf("iwlist %v scan | grep ESSID", a))
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	flag.Parse()

	s := fmt.Sprint(out)
	ss := string(s)
	sf := strings.Fields(ss)

	regex := regexp.MustCompile(`[123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz]{26,35}`)

	ssid1l := strings.TrimPrefix(sf[0], "ESSID:\"")
	ssid1 := strings.TrimSuffix(ssid1l, "\"")
	ssid2l := strings.TrimPrefix(sf[1], "ESSID:\"")
        ssid2 := strings.TrimSuffix(ssid2l, "\"")
	ssid3l := strings.TrimPrefix(sf[2], "ESSID:\"")
        ssid3 := strings.TrimSuffix(ssid3l, "\"")
	ssid4l := strings.TrimPrefix(sf[3], "ESSID:\"")
        ssid4 := strings.TrimSuffix(ssid4l, "\"")
	ssid5l := strings.TrimPrefix(sf[4], "ESSID:\"")
        ssid5 := strings.TrimSuffix(ssid5l, "\"")

	//Take address from cmd argument
        //arg := os.Args[1]

        //Handle debug args
        //arg1 := ""
        //if os.Args != nil && len(os.Args) > 1 {
                //arg1 = os.Args[1]
        //}


	if arg2 == "debug" {
		fmt.Println("SSID 1:", ssid1, "Length:", len(ssid1))
		fmt.Println("SSID 2:", ssid2, "Length:", len(ssid2))
		fmt.Println("SSID 3:", ssid3, "Length:", len(ssid3))
		fmt.Println("SSID 4:", ssid4, "Length:", len(ssid4))
		fmt.Println("SSID 5:", ssid5, "Length:", len(ssid5))
	}

	//fmt.Println("ssid 1 length: ", len(ssid1))
	//fmt.Println("\033[32mScanning for Skywire networks\033[0m")
	time.Sleep(4 * time.Second)

	if len(ssid1) == 32 && regex.MatchString(ssid1) == true {
		fmt.Println(ssid1)
	}
	//fmt.Println("\033[0mssid 2 length: ", len(ssid2))
	if len(ssid2) == 32 && regex.MatchString(ssid2) == true {
		fmt.Println(ssid2)
	}
	//fmt.Println("\033[0mssid 3 length: ", len(ssid3))
	if len(ssid3) == 32 && regex.MatchString(ssid3) == true {
		fmt.Println(ssid3)
	}
	//fmt.Println("\033[0mssid 4 length: ", len(ssid4))
	if len(ssid4) == 32 && regex.MatchString(ssid4) == true {
		fmt.Println(ssid4)
	}
	//fmt.Println("\033[0mssid 5 length: ", len(ssid5))
	if len(ssid5) == 32 && regex.MatchString(ssid5) == true {
		fmt.Println(ssid5)
		//fmt.Println("\033[0m")
	} /*
	   fmt.Println("ssid 6 length: ", len(sf[6]))
	   if len(sf[6])==40 {
	       fmt.Println("\033[36m6. \033[32m", sf[6])
	   }
	   fmt.Println("\033[0mssid 7 length: ", len(sf[7]))
	   if len(sf[7])==40 {
	       fmt.Println("\033[36m7. \033[32m", sf[7])
	   }
	   fmt.Println("\033[0mssid 8 length: ", len(sf[8]))
	   if len(sf[8])==40 {
	       fmt.Println("\033[36m8. \033[32m", sf[8])
	   }
	   fmt.Println("\033[0mssid 9 length: ", len(sf[9]))
	   if len(sf[9])==40 {
	       fmt.Println("\033[36m9. \033[32m", sf[9])
	   }
	   fmt.Println("\033[0mssid 10 length: ", len(sf[10]))
	   if len(sf[10])==40 {
	       fmt.Println("\033[36m10. \033[32m", sf[10])
	       fmt.Println("\033[0m")
	   }*/
//fmt.Println("\033[0m")
}
