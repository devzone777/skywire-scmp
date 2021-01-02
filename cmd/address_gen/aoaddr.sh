#!/bin/bash
until [[ $lent = 33 ]]
do
  
  go run address_gen.go > aoaddr.txt
  cat ./aoaddr.txt
  aoaddr=$( cat ./aoaddr.txt | grep address )
  aolen=${#aoaddr}
  lent=$[$aolen - 26]
  echo "$aoaddr"
  echo "$lent"
done
