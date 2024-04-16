CAST=${CAST:-cast}

THIS_ZSH="$0:A"
THIS_BASH="$BASH_SOURCE"
THIS=${BASH_SOURCE:-"$THIS_ZSH"}
REPO=$(dirname $(dirname "$THIS"))

DEVNET_D="${REPO}/.devnet"

DEVNET_FAUCET_ADDR=$(cast wallet derive-private-key 'test test test test test test test test test test test junk' | grep 'Address: ' | awk '{print $2}')
DEVNET_FAUCET_PRIV=$(cast wallet derive-private-key 'test test test test test test test test test test test junk' | grep 'Private key: ' | awk '{print $3}')

if [ ! -d "$DEVNET_D" ]; then
    echo "devnet is not initialized"
    exit 1
fi

export DEVNET_D=$DEVNET_D
export L1_RPC_URL="http://$(docker inspect ops-bedrock-l1-1  | jq -r '.[].NetworkSettings.Networks."ops-bedrock_default".IPAddress'):8545"
export L2_RPC_URL="http://$(docker inspect ops-bedrock-l2-1  | jq -r '.[].NetworkSettings.Networks."ops-bedrock_default".IPAddress'):8545"
# export INDEXER_URL="http://$(docker inspect indexer-index-1 | jq -r '.[].NetworkSettings.Networks."indexer_default".IPAddress'):8100"

l1-cast() {
    (
        set -xu
        ETH_RPC_URL=$L1_RPC_URL \
            $CAST "$@"
    )
}
l2-cast() {
    (
        set -xu
        ETH_RPC_URL=$L2_RPC_URL \
            $CAST "$@"
    )
}

devnet-addr() {
    local type=$1;
    cat "$DEVNET_D/addresses.json" | jq -r ".$type"
}

priv-for() {
    echo -n "$@" | sha3-256sum | awk '{print "0x" $1}'
}
addr-for() {
    $CAST wallet address $(priv-for "$@")
}

devnet-faucet() {
    local recipient=$1; shift
    local amount=$1;    shift
    local recipient_addr=$(addr-for "$recipient")

    local async=''
    local nonce=${NONCE:-1}
    local confirmations=${CONFIRMATIONS:-1}

    if [ ${ASYNC:-0} = 1 ]; then
        async='--async'
    fi


    l1-cast send \
        $async \
        --nonce $nonce \
        --confirmations "$confirmations" \
        --private-key "$DEVNET_FAUCET_PRIV" \
        --value "$amount" \
        "$recipient_addr"
}

devnet-l2-deposit() {
    local who=$1;       shift
    local amount=$1;    shift

    local async=''
    local confirmations=${CONFIRMATIONS:-1}

    if [ ${ASYNC:-0} = 1 ]; then
        async='--async'
    fi

    l1-cast send \
        $async \
        --confirmations "$confirmations" \
        --private-key "$(priv-for "$who")" \
        --value "$amount" \
        $(devnet-addr L1StandardBridgeProxy)
}

export L1_TO_L2_PROXY=$(devnet-addr L1StandardBridgeProxy)

