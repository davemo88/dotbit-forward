#!/bin/bash

if [[ -z $DATADIR || -z $RPCUSER || -z $RPCPASSWORD ]]; then
  echo "set all env: DATADIR RPCUSER RPCPASSWORD"
  exit 1
fi

RPCALLOWIP=$(getent ahosts ws | tail -1 | awk '{ print $1 }')

nmcconf=/home/nmcd/namecoin.conf
if [ -f $nmcconf ]; then
  rm $nmcconf
fi
mkdir -p $DATADIR
echo "rpcuser=$RPCUSER" >> $nmcconf
echo "rpcpassword=$RPCPASSWORD" >> $nmcconf
echo "rpcport=$RPCPORT" >> $nmcconf
echo "rpcallowip=$RPCALLOWIP" >> $nmcconf

namecoind -datadir=$DATADIR -conf=$nmcconf
