package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// "github.com/ubiq/go-ubiq/rlp"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// same private key is used
	// its a private key of the local ganache network
	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// recipient address
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0]) //<-- ts.GetRlp undefined
	// rawTxHex := hex.EncodeToString(rawTxBytes)
	// fmt.Printf(rawTxHex)
	// rawTxBytes := ts.GetRlp(0)
	rawTxHex := hex.EncodeToString(rawTxBytes)

	fmt.Println(rawTxHex) // f86...772
}
