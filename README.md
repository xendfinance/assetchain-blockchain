# xend-blockchain

EVM-compatible chain secured by the Lachesis consensus algorithm.

## Building the source

Building `xend-blockchain` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make xend-blockchain
```
The build output is ```build/xend-blockchain``` executable.

## Running `xend-blockchain`

Going through all the possible command line flags is out of scope here,
but we've enumerated a few common parameter combos to get you up to speed quickly
on how you can run your own `xend-blockchain` instance.

### Launching a network

Launching `xend-blockchain` readonly (non-validator) node for network specified by the genesis file:

```shell
$ xend-blockchain --genesis file.g
```

### Configuration

As an alternative to passing the numerous flags to the `xend-blockchain` binary, you can also pass a
configuration file via:

```shell
$ xend-blockchain --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to
export your existing configuration:

```shell
$ xend-blockchain --your-favourite-flags dumpconfig
```

#### Validator

New validator private key may be created with `xend-blockchain validator new` command.

To launch a validator, you have to use `--validator.id` and `--validator.pubkey` flags to enable events emitter.

```shell
$ xend-blockchain --nousb --validator.id YOUR_ID --validator.pubkey 0xYOUR_PUBKEY
```

`xend-blockchain` will prompt you for a password to decrypt your validator private key. Optionally, you can
specify password with a file using `--validator.password` flag.

#### Participation in discovery

Optionally you can specify your public IP to straighten connectivity of the network.
Ensure your TCP/UDP p2p port (5050 by default) isn't blocked by your firewall.

```shell
$ xend-blockchain --nat extip:1.2.3.4
```

## Dev

### Running testnet

The network is specified only by its genesis file, so running a testnet node is equivalent to
using a testnet genesis file instead of a mainnet genesis file:
```shell
$ xend-blockchain --genesis /path/to/testnet.g # launch node
```

It may be convenient to use a separate datadir for your testnet node to avoid collisions with other networks:
```shell
$ xend-blockchain --genesis /path/to/testnet.g --datadir /path/to/datadir # launch node
$ xend-blockchain --datadir /path/to/datadir account new # create new account
$ xend-blockchain --datadir /path/to/datadir attach # attach to IPC
```

### Testing

Lachesis has extensive unit-testing. Use the Go tool to run tests:
```shell
go test ./...
```

If everything goes well, it should output something along these lines:
```

### xend-blockchainting a private network (fakenet)

Fakenet is a private network optimized for your private testing.
It'll generate a genesis containing N validators with equal stakes.
To launch a validator in this network, all you need to do is specify a validator ID you're willing to launch.

Pay attention that validator's private keys are deterministically generated in this network, so you must use it only for private testing.

Maintaining your own private network is more involved as a lot of configurations taken for
granted in the official networks need to be manually set up.

To run the fakenet with just one validator (which will work practically as a PoA blockchain), use:
```shell
$ xend-blockchain --fakenet 1/1
```

To run the fakenet with 5 validators, run the command for each validator:
```shell
$ xend-blockchain --fakenet 1/5 # first node, use 2/5 for second node
```

If you have to launch a non-validator node in fakenet, use 0 as ID:
```shell
$ xend-blockchain --fakenet 0/5
```

After that, you have to connect your nodes. Either connect them statically or specify a bootnode:
```shell
$ xend-blockchain --fakenet 1/5 --bootnodes "enode://verylonghex@1.2.3.4:5050"
```

### Running the demo

For the testing purposes, the full demo may be launched using:
```shell
cd demo/
./start.sh # start the xend-blockchain processes
./stop.sh # stop the demo
./clean.sh # erase the chain data
```
Check README.md in the demo directory for more information.
