#!/bin/bash

go run ~/skywire-scmp/cmd/skywire-connect.go Dx5nqQpRidueMTquwgornCrrJHatqWKp & pkill wpa_supplicant & screen -dmS skywire-wpa-supplicant wpa_supplicant -c /etc/wpa_supplicant/wpa_supplicant.conf -i wlan0
