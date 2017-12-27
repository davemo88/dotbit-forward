#!/bin/bash
if [[ -z $DATADIR || -z $RPCUSER || -z $RPCPASSWORD || -z $RPCALLOWIP ]]; then
  echo "set all env: DATADIR RPCUSER RPCPASSWORD RPCALLOWIP"
  exit 1
fi

mkdir -p $DATADIR
nmcconf=/home/nmcd/namecoin.conf
echo "rpcuser=$RPCUSER" >> $nmcconf
echo "rpcpassword=$RPCPASSWORD" >> $nmcconf
echo "rpcport=$RPCPORT" >> $nmcconf
echo "rpcallowip=$RPCALLOWIP" >> $nmcconf

namecoind -datadir=$DATADIR -conf=$nmcconf
