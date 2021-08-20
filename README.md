<div align="center">
  <img src="https://public-luv.s3.us-east-2.amazonaws.com/Luv.svg">
</div>

---

Official node implementation of the [Avalanche](https://avax.network) network -
a blockchains platform with high throughput, and blazing fast transactions.

## Installation

Avalanche is an incredibly lightweight protocol, so the minimum computer requirements are quite modest.
Note that as network usage increases, hardware requirements may change.

- Hardware: 2 GHz or faster CPU, 6 GB RAM, >= 200 GB storage.
- OS: Ubuntu >= 18.04 or Mac OS X >= Catalina.
- Network: IPv4 or IPv6 network connection, with an open public port.
- Software Dependencies:
  - [Go](https://golang.org/doc/install) version >= 1.15.5 and set up [`$GOPATH`](https://github.com/golang/go/wiki/SettingGOPATH).
  - [gcc](https://gcc.gnu.org/)

### Native Install

Clone the Luv-Go repository in a different folder than the GOPATH, ie. Documents/ :

```sh
git clone https://github.com/hellobuild/Luv-Go.git
```
```sh
cd Luv-Go/
```
```sh
git checkout v1.4.9
```
```sh
./scripts/build.sh
```
```sh
./build/avalanchego --config-file=staking/development/node1.json
```
In another terminal, run a second development node

```sh
./build/avalanchego --config-file=staking/development/node2.json
```

To get a feeling of basic functions, import balance and others task, follow the commands in staking/desarrollo/flow.json in the same order. 


### Docker Install

Make sure docker is installed on the machine - so commands like `docker run` etc. are available.

Building the docker image of latest avalanchego branch can be done by running:

```sh
./scripts/build_image.sh
```

To check the built image, run:

```sh
docker image ls
```

The image should be tagged as `avaplatform/avalanchego:xxxxxxxx`, where `xxxxxxxx` is the shortened commit of the Avalanche source it was built from. To run the avalanche node, run:

```sh
docker run -ti -p 9650:9650 -p 9651:9651 avaplatform/avalanchego:xxxxxxxx /avalanchego/build/avalanchego
```

## Running Avalanche

### Connecting to Mainnet

To connect to the Avalanche Mainnet, run:

```sh
./build/avalanchego
```

You should see some pretty ASCII art and log messages.

You can use `Ctrl+C` to kill the node.

### Connecting to Fuji

To connect to the Fuji Testnet, run:

```sh
./build/avalanchego --network-id=fuji
```

### Creating a Local Testnet

To create a single node testnet, run:

```sh
./build/avalanchego --network-id=local --staking-enabled=false --snow-sample-size=1 --snow-quorum-size=1
```

This launches an Avalanche network with one node.

### Running protobuf codegen

To regenerate the protobuf go code, run `scripts/protobuf_codegen.sh` from the root of the repo

This should only be necessary when upgrading protobuf versions or modifying .proto definition files

To use this script, you must have [protoc](https://grpc.io/docs/protoc-installation/) and protoc-gen-go installed. protoc must be on your $PATH.

If you extract protoc to ~/software/protobuf/, the following should work:

```sh
export PATH=$PATH:~/software/protobuf/bin/:~/go/bin
go get google.golang.org/protobuf/cmd/protoc-gen-go
scripts/protobuf_codegen.sh
```
