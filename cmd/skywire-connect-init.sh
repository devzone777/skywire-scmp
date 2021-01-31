#!/bin/bash

go run ~/skywire-scmp/cmd/skywire-connect.go "$1" & pkill wpa_supplicant & screen -dmS skywire-wpa-supplicant wpa_supplicant -c /etc/wpa_supplicant/wpa_supplicant.conf -i wlan0
