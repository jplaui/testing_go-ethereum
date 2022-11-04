package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f872808504a817c800825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a764000080860308849c9ce8a05ae16086c37c8d489139305105a4968e86367b314884aa706b9fccda891c43ada00190e57279653b9d42a22972400fc3ab4b261fffe5bf104eb1642d5c13c4583e"
	// rawTx := "f86d8202b28477359400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a7640000802ca05924bde7ef10aa88db9c66dd4f5fb16b46dff2319b9968be983118b57bb50562a001b24b31010004f13d9a26b320845257a6cfc2bf819a3d55e3fc86263c5f0772"

	rawTxBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		log.Fatal("hex.DecodeString error:", err)
	}

	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		log.Fatal("rlp.DecodeBytes error:", err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal("client.SendTransaction error:", err)
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
