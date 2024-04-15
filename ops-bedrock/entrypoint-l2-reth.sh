#!/bin/sh
set -exu

ENABLE_TSHARK=${ENABLE_TSHARK:-0}

if [ "$ENABLE_TSHARK" = 1 ]; then
	tshark \
		-ni eth0 \
		-w "/tshark/$(date +'%Y%m%d-%H%M%S-')$(hostname)-L2-reth.pcapng" \
		-f 'tcp port 8545 or tcp port 8546 or tcp port 8551' &
fi


VERBOSITY=${GETH_VERBOSITY:-3}
GETH_DATA_DIR=/db
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
CHAIN_ID=$(cat "$GENESIS_FILE_PATH" | jq -r .config.chainId)
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"

if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
	echo "$GETH_CHAINDATA_DIR missing, running init"
	echo "Initializing genesis."
	op-reth init \
		--datadir="$GETH_DATA_DIR" \
		--chain="$GENESIS_FILE_PATH"
else
	echo "$GETH_CHAINDATA_DIR exists."
fi

# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the devnet.



exec op-reth node \
    --datadir="$GETH_DATA_DIR" \
	--http \
	--http.corsdomain="*" \
	--http.addr=0.0.0.0 \
	--http.port="$RPC_PORT" \
	--http.api=all \
	--ws \
	--ws.addr=0.0.0.0 \
	--ws.port="$WS_PORT" \
	--ws.origins="*" \
	--ws.api=all \
	--disable-discovery \
	--max-outbound-peers=0 \
    --max-inbound-peers=0 \
	--authrpc.addr="0.0.0.0" \
	--authrpc.port="8551" \
	--metrics 0.0.0.0:6060 \
    --chain="${GENESIS_FILE_PATH}" \
	--rollup.disable-tx-pool-gossip \
	--rollup.enable-genesis-walkback \
	"$@"

# --networkid="$CHAIN_ID" \
# --rpc.allow-unprotected-txs \

## error: the argument '--authrpc.jwtsecret <PATH>' cannot be used multiple times
# --authrpc.jwtsecret=/config/jwt-secret.txt \


## >> By default, Reth runs as an archive node. Such nodes have all historical blocks and the state at each of these blocks available for querying and tracing.
# --gcmode=archive \