package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("transaction hash hex:", tx.Hash().Hex())         // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("transaction value:", tx.Value().String())        // 10000000000000000
		fmt.Println("transaction gas:", tx.Gas())                     // 105000
		fmt.Println("transaction gas price:", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("transaction nonce:", tx.Nonce())                 // 110644
		fmt.Println("transaction data:", tx.Data())                   // []
		fmt.Println("transaction to hex", tx.To().Hex())              // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		// chainID, err := client.NetworkID(context.Background())
		// if err != nil {
		// log.Fatal("client.networkID", err)
		// }

		if msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), block.BaseFee()); err == nil {
			fmt.Println("transaction latest signer address in hex:", msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("receipt stats (1 success, 0 fail):", receipt.Status) // 1
	}

	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("transaction hash hex in block:", tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}

	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("transaction by hash hex:", tx.Hash().Hex())   // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println("check if transaction is pending:", isPending) // false
}
