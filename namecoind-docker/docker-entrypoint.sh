#!/bin/bash
set -e

if [ "$1" = 'namecoin' ]; then
    exec /data/namecoin/namecoin/bin/namecoind -datadir=/data -conf=/root/.namecoin/namecoin.conf `echo ${@:2}`
else
    exec "$@"
fi
