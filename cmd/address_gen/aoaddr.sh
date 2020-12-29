#!/bin/bash
#aolen=0
#aoaddr=""

#until [[ $aolen = 58 ]]
until [[ $lent = 32 ]]
do
  
  go run address_gen.go > aoaddr.txt
  cat /usr/local/go/github.com/skycoin/skycoin/cmd/address_gen/aoaddr.txt
  aoaddr=$( cat /usr/local/go/github.com/skycoin/skycoin/cmd/address_gen/aoaddr.txt | grep address )
  aolen=${#aoaddr}
  lent=$[$aolen - 26]
  echo "$aoaddr"
  #echo "$aolen"
  echo "$lent"
#  name=aoaddr.txt
#if [[ -e $name.ext || -L $name.ext ]] ; then
#    i=0
#    while [[ -e $name-$i.ext || -L $name-$i.ext ]] ; do
#        let i++
#    done
#    name=$name-$i
#fi
#touch -- "$name".ext
done
