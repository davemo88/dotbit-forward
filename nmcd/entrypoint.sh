#!/bin/bash
mkdir -p $DATADIR
nmcconf=$DATADIR/namecoin.conf
echo "rpcuser=$RPCUSER" >> $nmcconf
echo "rpcpassword=$RPCPASSWORD" >> $nmcconf
echo "rpcport=$RPCPORT" >> $nmcconf
echo "rpcallowip=$RPCALLOWIP" >> $nmcconf
#echo "rpcbind=$RPCBIND" >> $nmcconf

namecoind -datadir=$DATADIR -conf=$nmcconf
