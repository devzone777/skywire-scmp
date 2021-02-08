#!/bin/bash

swssid=$(go run skywire-scan.go)

go run ~/skywire-scmp/cmd/skywire-connect.go "$swssid" & pkill wpa_supplicant & screen -dmS skywire-wpa-supplicant wpa_supplicant -c /etc/wpa_supplicant/wpa_supplicant.conf -i wlan0
