#!/bin/sh
set -exu

ENABLE_TSHARK=${ENABLE_TSHARK:-0}

if [ "$ENABLE_TSHARK" = 1 ]; then
	apk add --no-cache tshark
	tshark \
		-ni eth0 \
		-w "/tshark/$(date +'%Y%m%d-%H%M%S-')$(hostname)-L2-op-node.pcapng" \
		-f 'tcp port 8545 or tcp port 8546 or tcp port 8551' &
fi

exec "$@"