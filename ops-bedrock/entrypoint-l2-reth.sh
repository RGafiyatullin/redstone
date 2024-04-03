#!/bin/sh
set -exu


RETH_DATA_DIR=/db
RETH_DATA_TOML="$RETH_DATA_DIR/reth.toml"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
CHAIN_ID=$(cat "$GENESIS_FILE_PATH" | jq -r .config.chainId)
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"
AUTH_RPC_PORT="${AUTH_RPC_PORT:-8551}"

if [ ! -f "$RETH_DATA_TOML" ]; then
	echo "$RETH_DATA_TOML missing, running init"
	echo "Initializing genesis."
	op-reth init \
		--datadir="$RETH_DATA_DIR" \
		--chain="$GENESIS_FILE_PATH" || exit $?
else
	echo "$RETH_DATA_TOML exists."
fi

# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the devnet.

exec op-reth node \
    --datadir="$RETH_DATA_DIR" \
	--http \
	--http.corsdomain="*" \
	--http.addr=0.0.0.0 \
	--http.port="$RPC_PORT" \
	--http.api=web3,debug,eth,txpool,net \
	--ws \
	--ws.addr=0.0.0.0 \
	--ws.port="$WS_PORT" \
	--ws.origins="*" \
	--ws.api=debug,eth,txpool,net \
	--disable-discovery \
	--max-outbound-peers=0 \
    --max-inbound-peers=0 \
	--authrpc.addr="0.0.0.0" \
	--authrpc.port="${AUTH_RPC_PORT}" \
	--metrics 0.0.0.0:6060 \
    --chain="${GENESIS_FILE_PATH}" \
	"$@"

# --networkid="$CHAIN_ID" \
# --rpc.allow-unprotected-txs \

## error: the argument '--authrpc.jwtsecret <PATH>' cannot be used multiple times
# --authrpc.jwtsecret=/config/jwt-secret.txt \


## >> By default, Reth runs as an archive node. Such nodes have all historical blocks and the state at each of these blocks available for querying and tracing.
# --gcmode=archive \