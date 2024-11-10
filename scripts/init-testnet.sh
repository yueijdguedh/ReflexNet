#!/bin/bash

# Initialize a local testnet for ReflexNet

BINARY=reflexnetd
CHAIN_ID=reflexnet-testnet-1
MONIKER="reflexnet-node"
KEYRING="test"
KEYALGO="secp256k1"

# remove existing daemon
rm -rf ~/.reflexnet

# Build binary
make build

# Initialize chain
$BINARY init $MONIKER --chain-id $CHAIN_ID

# Add test keys
$BINARY keys add alice --keyring-backend $KEYRING --algo $KEYALGO
$BINARY keys add bob --keyring-backend $KEYRING --algo $KEYALGO

# Add genesis accounts
$BINARY genesis add-genesis-account alice 100000000000mcell --keyring-backend $KEYRING
$BINARY genesis add-genesis-account bob 100000000000mcell --keyring-backend $KEYRING

# Create gentx
$BINARY genesis gentx alice 10000000mcell --chain-id $CHAIN_ID --keyring-backend $KEYRING

# Collect gentxs
$BINARY genesis collect-gentxs

echo "Testnet initialized successfully!"
echo "Start the node with: $BINARY start"

