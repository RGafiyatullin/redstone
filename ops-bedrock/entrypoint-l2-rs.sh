#!/bin/sh
set -exu

ENABLE_TSHARK=${ENABLE_TSHARK:-0}

if [ "$ENABLE_TSHARK" = 1 ]; then
	tshark \
		-ni eth0 \
		-w "/tshark/$(date +'%Y%m%d-%H%M%S-')$(hostname)-L2-rs.pcapng" \
		-f 'tcp port 8545 or tcp port 8546 or tcp port 8551' &
fi

OP_RETH_BIN=${OP_RETH_BIN:-$(which op-reth)}

VERBOSITY=${GETH_VERBOSITY:-3}
GETH_DATA_DIR=/db
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
FORK_CONDITIONS_FILE_PATH=${FORK_CONDITIONS_FILE_PATH:-/fork-conditions.json}
RETH_CHAIN_FILE_PATH="/tmp/reth-chain.json"

jq \
	--slurpfile fork_conditions "${FORK_CONDITIONS_FILE_PATH}" \
	'{
		"chain": .config.chainId,
		"genesis": .,
		"hardforks": $fork_conditions[0],
		"base_fee_params": {
			"max_change_denominator": "0x8",
			"elasticity_multiplier": "0x2"
		}
	}' \
	< "$GENESIS_FILE_PATH" \
	> "$RETH_CHAIN_FILE_PATH"

CHAIN_ID=$(cat "$GENESIS_FILE_PATH" | jq -r .config.chainId)
RPC_PORT="${RPC_PORT:-8545}"
WS_PORT="${WS_PORT:-8546}"

if [ ! -d "$GETH_CHAINDATA_DIR" ]; then
	echo "$GETH_CHAINDATA_DIR missing, running init"
	echo "Initializing genesis."
	${OP_RETH_BIN} init \
		--datadir="$GETH_DATA_DIR" \
		--chain="$RETH_CHAIN_FILE_PATH"
else
	echo "$GETH_CHAINDATA_DIR exists."
fi

# Warning: Archive mode is required, otherwise old trie nodes will be
# pruned within minutes of starting the devnet.


exec "${OP_RETH_BIN}" node \
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
    --chain="${RETH_CHAIN_FILE_PATH}" \
	--rollup.disable-tx-pool-gossip \
	--rollup.enable-genesis-walkback \
	"$@"
