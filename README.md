## Repository for Ethereum development in Golang

### Goethereumbook
The link to the documentation can be found here: https://goethereumbook.org/en/ and https://github.com/miguelmota/ethereum-development-with-go-book
Below, you find the documentation on how scripts can be executed per topic. Always make sure to jump into the respective folders before executing the commands. e.g. jump into the `client` folder before runnnig and then execute commands like `go run main.go`.

This time, system dependencies relied on:
```
$$$ solc --version
solc, the solidity compiler commandline interface
Version: 0.8.17+commit.8df45f5f.Linux.g++
```
```
$$$ ganache-cli --version
Ganache CLI v6.12.2 (ganache-core: 2.13.2)
```
```
$$$ go version
go version go1.19.3 linux/amd64
```
```
$$$ geth version
Geth
Version: 1.10.25-stable
Git Commit: 69568c554880b3567bace64f8848ff1be27d084d
Git Commit Date: 20220915
Architecture: amd64
Go Version: go1.19.1
Operating System: linux
GOPATH=/home/jp/Documents/coding/go
GOROOT=
```


#### client folder
- install ganache with `npm install -g ganache-cli`, check with `ganache-cli`
- start ganache with same sequence of public addresses `ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"`, run this command in another terminal
- inside the `client` folder, run `go mod tidy`, next run `go run ganache_client.go`
- if you do not start ganache, you can run `go run remote_client.go`

#### accounts folder
- init the folder with `go mod tidy`
- execute `address.go` with `go run address.go`
- execute `wallet_generate.go` with `go run wallet_generate.go`
- execute `account_balance.go` with `go run account_balance.go` (make sure to start ganache with the above mentioned mnemonic phrase)
- execute `address_check.go` with `go run address_check.go`
- execute `keystore.go` with `go run keystore.go -create true` and `go run keystore.go -import true`

#### transactions folder
- init the folder with `go mod tidy`

#### smartcontracts folder

### Useful links
- Block explorer Etherscan [here](https://etherscan.io)
- File sotrage Swarm [here]()
- P2P messaging Whisper [here]()
- Ganache [here](https://trufflesuite.com/ganache/)
- Infura [here]()
-

### Running your own local network with geth

