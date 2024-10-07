# AssetChain Blockchain

This repository contains all resources required to set up a full node in the AssetChain Blockchain.

## Table of Contents
- [Requirements](#requirements)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)
- [Support](#support)

## Requirements

The requirements for running a full node are shown below:

1. **Minimum Hardware Requirements**: AWS EC2 m5.xlarge with 4 vCPUs (3.1 GHz) and at least 1 TB of Amazon EBS General Purpose SSD (gp2) storage (or equivalent).
2. **Recommended Specifications**: For better performance, consider using AWS m6i.2xlarge, c6i.4xlarge, or bare metal with equivalent or higher specs.

## Getting Started

To get started, follow the steps below:

1. **Launch Cloud Instance**:
   - You can either run a node on your own hardware or use a cloud provider. We recommend using one of the major cloud providers, such as Amazon AWS.
   - We recommend using Ubuntu Server 22.04 LTS (64-bit).

2. **Network Settings**:
   - Open up port 22 or any other port of preference for SSH.
   - Open up port 5050 for both TCP and UDP traffic. Alternatively, use the `--port <port>` flag to specify a custom port when running your node.

3. **Set up a Non-Root User**:
   - SSH into your machine:
     ```bash
     ssh root@{FULLNODE_IP_ADDRESS}
     ```
   - Update the system:
     ```bash
     sudo apt-get update && sudo apt-get upgrade -y
     ```
   - Create a non-root user:
     ```bash
     USER={USERNAME}
     sudo mkdir -p /home/$USER/.ssh
     sudo touch /home/$USER/.ssh/authorized_keys
     sudo useradd -d /home/$USER $USER
     sudo usermod -aG sudo $USER
     sudo chown -R $USER:$USER /home/$USER/
     sudo chmod 700 /home/$USER/.ssh
     sudo chmod 644 /home/$USER/.ssh/authorized_keys
     ```
   - Add your SSH public key to the new user's `authorized_keys` file.

4. **Install Required Tools**:
   - Install build tools:
     ```bash
     sudo apt install apt-transport-https libc-bin gcc git musl ca-certificates curl zsh make unzip build-essential gawk wget bison flex
     ```
   - Install Go:
     ```bash
     wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
     sudo tar -xvf go1.22.0.linux-amd64.tar.gz
     sudo mv go /usr/local
     export GOROOT=/usr/local/go
     export GOPATH=$HOME/go
     export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
     source ~/.profile
     exec bash
     ```

5. **Install and Set Up Opera**:
   - Clone the Opera repository and build:
     ```bash
     git clone https://github.com/xendfinance/assetchain-blockchain
     cd assetchain-blockchain/
     go clean -modcache
     go mod tidy
     export GOPROXY="https://proxy.golang.org"
     mkdir -pv build
     go build -v -o build/opera ./cmd/opera
     ```

6. **Start Opera Node**:
   - Download the genesis file and start the node:
     ```bash
     curl https://asset-testnet.assetchain.org/genesis_new.g --output genesis.g
     cd build/
     nohup ./opera --port 3000 --nat any --genesis.allowExperimental --genesis ../genesis.g --http --http.addr="0.0.0.0" --http.port=4000 --http.corsdomain=* --http.vhosts=* --http.api=eth,debug,net,admin,web3,personal,txpool,ftm,dag --bootnodes="enode://27c5f90bd11d2e5df3901c8f893cfcbe0e62c0edfda88170eff43a87eb54c333a1ddce3dc6765eeeccfd37f01e614373e2d0449512735e4a96f528ea53e87ddf@34.147.162.187:3000" > opera.log &
     ```
     N:B if you only want to run fullnode for rpc then you can stop here.


## Contributing

See [CONTRIBUTING.md](https://github.com/xendfinance/nodesale/CONTRIBUTING.md) for contribution and pull request protocol. We expect contributors to follow our guide when submitting code or comments.

## License

[![License: GPL v3.0](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Support

For questions or suggestions, just say Hi on [Telegram](https://t.me/xendfinancedevs).  
We're always glad to help.
