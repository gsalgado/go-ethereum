#!/usr/bin/env bash

# Compile the solidity code.
solc -o $(pwd) --abi Proxy.sol && \
solc -o $(pwd) --bin Proxy.sol && \

# Generate the bindings for the contract.
abigen --pkg mock --type Proxy --abi Proxy.sol:Proxy.abi --bin Proxy.bin --out proxy.go && \

# Clean up the ABI and bin files.
rm $(pwd)/*.abi $(pwd)/*.bin
