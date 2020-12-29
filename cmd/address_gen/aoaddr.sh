#!/bin/bash
until [[ $lent = 32 ]]
do
  
  go run address_gen.go > aoaddr.txt
  cat /usr/local/go/github.com/skycoin/skycoin/cmd/address_gen/aoaddr.txt
  aoaddr=$( cat /usr/local/go/github.com/skycoin/skycoin/cmd/address_gen/aoaddr.txt | grep address )
  aolen=${#aoaddr}
  lent=$[$aolen - 26]
  echo "$aoaddr"
  echo "$lent"
done
