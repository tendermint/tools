#!/bin/sh

#BINARY = build/tendermint
#ID
#SIZE

BINARY_TYPE="`file $BINARY`"
if ! [[ "$BINARY_TYPE" =~ "ELF 64-bit LSB executable, x86-64" ]]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

export TMHOME=/tendermint${ID}${SIZE}
$BINARY $@

