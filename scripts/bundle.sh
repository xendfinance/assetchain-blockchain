#!/bin/bash

IFS=$'\n' read -d '' -r -a valPubKeys < ./validators.txt
N=${#valPubKeys[@]}

network="mainnet"
i=1

input="./validators.txt"

for pubkey in "${valPubKeys[@]}"
do
	datadir="$network/$i"
	echo "$datadir - START"
	mkdir -pv "$datadir"
	cp "./pass.txt" "$datadir/pass.txt"
	mv "datadir/datadir_opera$i" "$datadir/datadir"
	echo "$pubkey" > "$datadir/validators.txt"

	zip -r -9 "$network/$i".zip "$datadir"
	rm -rf "$datadir"

	((i++))
done

zip -r -9 "$network"/genesis.g.zip genesis.g

