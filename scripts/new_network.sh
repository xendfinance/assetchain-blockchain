#!/usr/bin/env bash

validutorsNum=3

../build/opera \
network new $validutorsNum \
--datadir ./datadir/datadir_opera \
--password ./pass.txt \
--validatorsfile ./validators.txt \
--network.type devnet

../build/opera --datadir ./datadir/datadir_opera1 export genesis genesis.g
