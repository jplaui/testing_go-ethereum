package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"

	"smartcontracts/api"
	"smartcontracts/store"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var flagDeploy = flag.Bool("deploy", false, "set go run main -deploy true to deploy a contract")
var flagDeployStore = flag.Bool("deploystore", false, "set -deploystore true to deploy a contract")
var flagLoadStore = flag.Bool("loadstore", false, "set -loadstore true to load a contract")
var flagQueryStore = flag.Bool("querystore", false, "set -querystore true to query a contract method")
var flagWriteStore = flag.Bool("writestore", false, "set -writestore true to write to the store contract")
var flagCodeStore = flag.Bool("codestore", false, "set -codestore true to get the bytecode of the store contract")

func main() {

	// parsing flags
	flag.Parse()
	if *flagDeploy {

		// address of etherum env
		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Println("ethclient.Dial error:", err)
			return
		}

		// create auth and transaction package for deploying smart contract
		auth := getAccountAuth(client, "f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")

		//deploying smart contract
		deployedContractAddress, tx, instance, err := api.DeployApi(auth, client) //api is redirected from api directory from our contract go file
		if err != nil {
			log.Println("api.DeployApi() error:", err)
			return
		}

		// print deployed contract address
		log.Println("deployed contract address:", deployedContractAddress.Hex())
		// print transaction hash in hex
		log.Println("deployed contract transaction hash in hex:", tx.Hash().Hex())

		_ = instance

		log.Println("instance:", instance)

	}

	if *flagDeployStore {

		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		input := "1.0"
		address, tx, instance, err := store.DeployStore(auth, client, input)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("contract address:", address.Hex())                // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
		fmt.Println("contract transaction hash hex:", tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

		_ = instance

	}

	if *flagLoadStore {

		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatal(err)
		}

		address := common.HexToAddress("0x4A31ECe693fB614935eFB337034F9C79efEC03B5")
		instance, err := store.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("contract is loaded")
		_ = instance

	}

	if *flagQueryStore {

		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatal(err)
		}

		address := common.HexToAddress("0x4A31ECe693fB614935eFB337034F9C79efEC03B5")
		instance, err := store.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		version, err := instance.Version(nil)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("contract version:", version) // "1.0"
	}

	if *flagWriteStore {

		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		address := common.HexToAddress("0x4A31ECe693fB614935eFB337034F9C79efEC03B5")
		instance, err := store.NewStore(address, client)
		if err != nil {
			log.Fatal(err)
		}

		key := [32]byte{}
		value := [32]byte{}
		copy(key[:], []byte("foo"))
		copy(value[:], []byte("bar"))

		tx, err := instance.SetItem(auth, key, value)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("setItem tx hash hex: %s", tx.Hash().Hex())

		result, err := instance.Items(nil, key)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("query key return value:", string(result[:])) // "bar"

	}

	if *flagCodeStore {

		client, err := ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatal(err)
		}

		contractAddress := common.HexToAddress("0x4A31ECe693fB614935eFB337034F9C79efEC03B5")
		bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
		if err != nil {
			log.Fatal(err)
		}

		log.Println("store contract bytecode:", hex.EncodeToString(bytecode)) // 60806...10029
	}

}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	log.Println("nonce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units (300000)
	auth.GasPrice = gasPrice

	return auth
}
