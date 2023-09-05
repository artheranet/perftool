# Arthera performance testing tool
Stress test and measure the performance of an Arthera network.

# Build instructions
Just run `make` to create the `build/txgen` binary.

Instructions for running:

1. Generate some fake accounts. This will create a folder called `keys_txgen` that will have the private keys of the generated fake accounts:
```shell
./build/txgen fakeaccs
```
2. Update the `txgen.toml` config file with:
    - `ChainId` - the chain id of the Arthera chain you want to connect to
    - `Payer` - the address of the Arthera account that will pay for the transactions
    - `URLs` - a list of WebSocket urls to connect to
3. Import the payer key from your node data dir to the local folder _keys_txgen_:
```shell
./build/txgen importacc 0x0c08A529D58152A01d20b46B28DEEB7a4075104A ~/.arthera/keystore "fakepassword"
```

5. Init fake account balances with funds from the payer account:
```shell
./build/txgen initbalance
```

6. Run the transfer TPS test:
```shell
./build/txgen transfers
```

7. Run the contract calls TPS test:
```shell
./build/txgen calls
```

The tool will print the average TPS while running, and it will run indefinitely until you stop it with Ctrl+C.
