#!/usr/bin/env bash

validutorsNum=5

#../build/opera \
#network new $validutorsNum \
#--datadir ./datadir/datadir_opera \
#--network.type=devnet \
#--password ./pass.txt \
#--validatorsfile ./validators.txt

../build/opera --datadir ./datadir/datadir_opera1 export genesis genesis.g
