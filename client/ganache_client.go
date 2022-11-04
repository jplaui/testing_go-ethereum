package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully connected client")
	_ = client

}
