#!/bin/sh
set -exu

ENABLE_TSHARK=${ENABLE_TSHARK:-0}

if [ "$ENABLE_TSHARK" = 1 ]; then
	tshark \
		-ni eth0 \
		-w "/tshark/$(date +'%Y%m%d-%H%M%S-')$(hostname)-L2-rs.pcapng" \
		-f 'tcp port 8545 or tcp port 8546 or tcp port 8551' &
fi

REDSTONE_SEQUENCER_BIN=${REDSTONE_SEQUENCER_BIN:-$(which redstone-sequencer-lite)}

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

exec "${REDSTONE_SEQUENCER_BIN}" node --chain-spec "$RETH_CHAIN_FILE_PATH" "$@"