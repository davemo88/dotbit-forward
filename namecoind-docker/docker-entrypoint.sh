#!/bin/bash
set -e

if [ "$1" = 'namecoin' ]; then
    echo "starting namecoin"
    exec /data/namecoin/namecoin/bin/namecoind -datadir=/namecoin_data -conf=/root/.namecoin/namecoin.conf `echo ${@:2}` 
else
    echo "passing command through"
    exec "$@"
fi
