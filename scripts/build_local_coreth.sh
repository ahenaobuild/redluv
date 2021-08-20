#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Directory above this script
AVALANCHE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the versions
source "$AVALANCHE_PATH"/scripts/versions.sh

# Load the constants
source "$AVALANCHE_PATH"/scripts/constants.sh

# Build Coreth
echo "Building Coreth @ ${coreth_version} ..."
cd "$AVALANCHE_PATH"
coreth_local_path="$AVALANCHE_PATH/coreth"
go build -ldflags "-X github.com/ava-labs/coreth/plugin/evm.Version=$git_commit" -o "$latest_evm_path" "$coreth_local_path/plugin/"*.go

# Building coreth + using go get can mess with the go.mod file.
go mod tidy

# INSTRUCTIONS TO GET CORETH LOCAL

#Create new directory coreth
#		from https://github.com/ava-labs/coreth copy the following directories:
#			accounts. chain, consensus, core, eth, interfaces, internal, miner, node, params, plugin, rpc, signer, tests

#Do the following replaces into the root coreth directory:
#		from "AVAX" to "LUV"
#		from github.com/ava-labs/avalanchego to github.com/hellobuild/Luv-Go
#		from github.com/ava-labs/coreth to github.com/hellobuild/Luv-Go/coreth
#		from AVAXAssetID to LUVAssetID
#		from MilliAvax to MilliLuv

#Download/Update dependencies and update go.mod / go.sum
#		go mod tidy
