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
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal("client dial", err)
	}

	blockNumber := big.NewInt(1)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("blockbynumber:", err)
	}

	for _, tx := range block.Transactions() {
		fmt.Println("transaction hash hex:", tx.Hash().Hex())         // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("transaction value:", tx.Value().String())        // 10000000000000000
		fmt.Println("transaction gas:", tx.Gas())                     // 105000
		fmt.Println("transaction has price:", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("transaction nonce:", tx.Nonce())                 // 110644
		fmt.Println("transaction data:", tx.Data())                   // []
		fmt.Println("transaction to hex address:", tx.To().Hex())     // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		// chainID, err := client.NetworkID(context.Background())
		// if err != nil {
		// log.Fatal(err)
		// }

		if msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), block.BaseFee()); err == nil {
			fmt.Println("transaction sender address hex:", msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal("transactionreceipt", err)
		}

		fmt.Println("transaction receipt (1 indicates success,  0 fail):", receipt.Status) // 1
	}

	blockHash := common.HexToHash("0xaa3e27f9c6f6c604fa37e50081b8f0e23e59c7fb87238d7ef34079186a6b7ac6")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal("transaction count for blockhash:", err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("tx in block hash hex:", tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}

	txHash := common.HexToHash("0xd678bd8b11c636b38221c402ff4b29e1ddfb9604ca57368cd9a4df4446582941")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal("transaction by hash error:", err)
	}

	fmt.Println("tx by hash hex:", tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println("tx is pending:", isPending)        // false
}
