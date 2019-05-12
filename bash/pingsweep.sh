#!/bin/bash

net_add=""
interface=""
mask=""

while test $# -gt 0; do
  case "$1" in
    -h|--help)
      echo "$0 [options]"
      echo " IMPORTANT!! only ONE option allowed"
      echo "options:"
      echo "-h, --help                show brief help"
      echo "-n, --net                 specify the net adress and it's mask (ej: 10.10.10.10/24)"
      echo "-i, --interface           specify the interface name to analize"
      echo "if no options used it will analize the default interface"
      exit 0
      ;;
    -n|--net)
      shift
      if test $# -gt 0; then
        net_add=$1
      else
        echo "no net address given"
        exit 1
      fi
      shift
      ;;
    -i|--interface)
      shift
      if test $# -gt 0; then
        interface=$1
      else
        echo "no interface name given"
        exit 1
      fi
      shift
      ;;
    *)
      break
      ;;
  esac
done

ip2net() {
  IP=$(echo $1 | cut -d / -f 1)
  PREFIX=$(echo $1 | cut -d / -f 2)

  IFS=. read -r i1 i2 i3 i4 <<< $IP
  IFS=. read -r xx m1 m2 m3 m4 <<< $(for a in $(seq 1 32); do
    if [ $(((a - 1) % 8)) -eq 0 ]; then
      echo -n .;
    fi;
    if [ $a -le $PREFIX ]; then
      echo -n 1;
    else
      echo -n 0;
    fi;
  done)
  printf "%d.%d.%d.%d\n" "$((i1 & (2#$m1)))" "$((i2 & (2#$m2)))" "$((i3 & (2#$m3)))" "$((i4 & (2#$m4)))"
}

getMask() {
  echo $1 | cut -d / -f 2
}

getNetAddrFromInt() {
  line=$(ip -o -f inet addr show $1)

  net=$(echo $line | cut -d " " -f 4)

  net_add=$(ip2net $net)
  mask=$(getMask $net)
}

# default case
if [ -z "$net_add" ] && [ -z "$interface" ]; then
  echo "default"
  def_int="$(ip route list | grep default | cut -d ' ' -f 5)"
  getNetAddrFromInt $def_int

elif [ -n "$net_add" ]; then
  echo "net"
  mask=$(getMask $net_add)
  net_add=$(ip2net $net_add)

elif [ -n "$interface" ]; then
  echo  "interface"
  getNetAddrFromInt $interface

fi

echo $net_add $mask


n=$((2**(32-mask)-1))
for (( i=1; i<$n; i++ )); do
  IFS=. read -r i1 i2 i3 i4 <<< $net_add
  ip=$(printf "%d.%d.%d.%d\n" "$((i/16581375 > 0 ? i1 + ((i/16581375) %256) : i1))" "$((i/65025 > 0 ? i2 + ((i/65025) %256) : i2))" "$((i/256 > 0 ? i3 + ((i/256) %256) : i3))" "$((i4 + i%256))")

  ( ping -c 3 -t 5 $ip > /dev/null && echo $ip is Alive ) &
done;

for job in $(jobs -p); do
    wait $job
done

exit 0
