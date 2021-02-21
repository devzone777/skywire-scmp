#!/bin/bash

apt-update -y && apt-upgrade -y && apt-install -y screen hostapd git

DIRECTORY=/root/.scmp/swap
if [ -d "$DIRECTORY" ]; then
  exit 0
else
  if [ ! -d "$DIRECTORY" ]; then
    mkdir -p /root/.scmp/swap
fi

cd /root/skywire-scmp

skywire-cli visor gen-config && hypervisor gen-config

screen -dmS skywire_visor skywire-visor #&& screen -dmS HYPERVISOR hypervisor

./address_gen > /root/.scmp/swap/wallet.wlt

addr=$(cat /root/.scmp/swap/wallet.wlt | grep address | sed 's/^.* / /' | tr -d \" | tr -d \, | tr -d \ )
auxaddr=$(/root/skywire-scmp/auxaddr_gen $addr | grep Auxillary | sed 's/^.* / /' | tr -d \ )
psk=$(/root/skywire-scmp/skywire-gen-psk $auxaddr)

/root/skywire-scmp/skywire-hapd $auxaddr $psk > /etc/hostapd/hostapd.conf
echo DAEMON_CONF="/etc/hostapd/hostapd.conf" >> /etc/default/hostapd
systemctl unmask hostapd
systemctl enable hostapd
rfkill unblock wlan
systemctl start hostapd

apt install -y dnsmasq

DEBIAN_FRONTEND=noninteractive apt install -y netfilter-persistent iptables-persistent

echo -e "interface wlan0\n    static ip_address=10.10.0.1/24\n    nohook wpa_supplicant" > /etc/dhcpd.conf

touch /etc/sysctl.d/routed-ap.conf
echo -e net.ipv4.ip_forward=1 > /etc/sysctl.d/routed-ap.conf

iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE

netfilter-persistent save

mv /etc/dnsmasq.conf /etc/dnsmasq.conf.orig
touch /etc/dnsmasq.conf

echo -e "interface=wlan1\ndhcp-range=10.10.0.2,10.10.0.20,255.255.255.0,24h\ndomain=skywire\naddress=/skywire.ap/10.10.0.1" > /etc/dnsmasq.conf

screen -dmS swap /root/skywire-scmp/swap

scanresult=$(/root/skywire-scmp/skywire-scan)
/root/skywire-scmp/skywire-connect $scanresult
pkill wpa_supplicant
screen -dmS skywire-wpa-supplicant wpa_supplicant -c /etc/wpa_supplicant/wpa_supplicant.conf -i wlan0

wlan0ip=$(ip route show 0.0.0.0/0 dev wlan1 | cut -d\\  -f3)
rpk=$(go run /root/skywire-scmp/skywire-get-remote-ap-pk.go $wlan0ip)

skywire-cli visor add-tp $rpk

exit 0
