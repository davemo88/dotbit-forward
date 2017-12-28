#!/bin/bash

NMCD_IP=$(getent ahosts nmcd | tail -1 | awk '{ print $1 }')
export NMCD_HOST="http://$NMCD_IP:8336"

./ws
