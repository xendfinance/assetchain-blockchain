#!/usr/bin/env bash

validutorsNum=5

../build/opera \
network new $validutorsNum \
--datadir ./datadir/datadir_opera \
--password ./pass.txt \
--validatorsfile ./validators.txt \
--network.type=devnet
