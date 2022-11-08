## Repository for Ethereum development in Golang

### Goethereumbook
The link to the documentation can be found here: https://goethereumbook.org/en/ and https://github.com/miguelmota/ethereum-development-with-go-book

Below, you find the documentation on how scripts can be executed per topic. Always make sure to jump into the respective folders before executing the commands. e.g. jump into the `client` folder before executing commands like `go run main.go`.

This time, system dependencies rely on:
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
(arch install truffle by cloning the aur truffle [github](https://aur.archlinux.org/truffle.git), switch into folder an install with `makepkg -si`)
```
$$$ truffle
```
(arch install ganache-bin binary from aur [repo](https://aur.archlinux.org/ganache-bin.git), switch into cloned folder and run `makepkg -si` to install)
```
ganache binary version: ganache-2.5.4
```
Additionally, you can download the ganache linux appimage [here](https://trufflesuite.com/ganache/) to get [ganache-ui](https://github.com/trufflesuite/ganache-ui). After making the appimage executable by running `chmod +x ganache-2.5.4-linux-x86_64.AppImage`, you can start it with `./ganache-2.5.4-linux-x86_64.AppImage`. When using the appimage of ganache, make sure to use the url and accounts associated with the ganache network configured with appImage (to get the same accounts, paste in the mnemonic phrase when creating a new workspace). Ganache UI can be linked to a truffle work environment for better debugging when developing smart contracts...

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
- execute `blocks.go` with `go run blocks.go`
- execute `gblocks.go` (ganache blocks) with `go run gblocks.go`
- execute `transactions.go` with `go run transactions.go`
- (create blocks with transactions first in ganache) execute `gtransactions.go` with `go run gtransactions.go`. make sure to update the block hash and transaction hash and set a sufficiently low block number if you have not send many transactions (e.g. script works if you call `gtransaction_raw_create.go` and `gtransaction_raw_sendreate.go` first).
- execute `gtransaction_raw_create.go` (ganache, edit the file and make sure you are using a private key of an account registered in ganache genesis) with `go run gtransaction_raw_create.go`
- copy the output of `go run gtransaction_raw_create.go` and paste it into `rawTx` in the file `gtransaction_raw_sendreate.go` and execute `go run gtransaction_raw_sendreate.go`. be aware, if you call `go run gtransaction_raw_sendreate.go` another time, make sure to create new transaction bytes with a new nonce, otherwise, replayed nonce values from a same account will be rejected.
- (optional) you can query and inspect your transaction if you are using ganache
- execute `gtransfer_eth.go` with `go run gtransfer_eth.go`, make sure to update the private key and address of the recipient if you wanna see the eth transfer in ganache.
- TODO: execute `gtransfer_token.go` with TODO
- TODO: execute `block_subscribe.go` with `go run block_subscribe.go`

#### smartcontracts folder
- initially run `go mod tidy`
- compile solidity and create ABIs of contracts inside `contracts` folder with `solc --optimize --abi ./contracts/BalanceCheck.sol -o build`
- next, create the bin files with `solc --optimize --bin ./contracts/MySmartContract.sol -o build`
- create `api` folder with `mkdir api`
- next create the go file wrapper with `abigen --abi=./build/BalanceCheck.abi --bin=./build/BalanceCheck.bin --pkg=api --out=./api/BalanceCheck.go`
- next, by running main `go run main -deploy true` with flag for contract deployment returns output:
```
2022/11/08 12:23:15 deployed contract address: 0x0665C95eeA208ed16bdF9f50e762383FC6A77701
2022/11/08 12:23:15 deployed contract transaction hash in hex: 0x5a3e3d79d0bb8133f06a533bce96c1ad11b38db81056c05d188425d8397795e0
2022/11/08 12:23:15 instance: &{{0xc000005180} {0xc000005180} {0xc000005180}}
```
running it a second time returns:
```
2022/11/08 12:27:12 nonce= 3
2022/11/08 12:27:12 deployed contract address: 0x0e93A5BF00452E3b43Cc7641C8387d1AE6997f5B
2022/11/08 12:27:12 deployed contract transaction hash in hex: 0xfe14127ae2a41e9271ace699e494ecf223bc5f5461b413e21502010697e9f22b
2022/11/08 12:27:12 instance: &{{0xc0000f2280} {0xc0000f2280} {0xc0000f2280}}
```
running `go run main.go -deploystore true` returns:
```
contract address: 0x4A31ECe693fB614935eFB337034F9C79efEC03B5
contract transaction hash hex: 0x20c8d04883accec6f657bb40d33c49e1190e8ec6a5be4c4c9fcd947df92bc905
```
- run `go run main.go -loadstore true` loads the store contract
- run `go run main.go -querystore true` to query the store contract version
- to set/write a value into the store, use `go run main.go -writestore true`
- get the store contract bytecode with `go run main.go -codestore true`


#### events
- initially start with `go mod tidy` inside the folder
-


### Useful links
- Block explorer Etherscan [here](https://etherscan.io)
- File sotrage Swarm [here]()
- P2P messaging Whisper [here]()
- Ganache [here](https://trufflesuite.com/ganache/)
- Infura [here]()
-

### Running your own local network with geth

