package main

import (
	"credits-go/api"
	"credits-go/core"
	"credits-go/general"
	"credits-go/utils"
	"fmt"
)

var _ = general.GoUnusedProtection__
var _ = api.GoUnusedProtection__

func main() {

	//A Credits Node API End-Point
	addr := "localhost:9090"

	testGetWalletData(addr)

}

func testGetWalletData(node string) {

	//generate new Credits wallet using utility feature
	publicKey, privateKey, err := utils.GenerateKeys()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public Key: %s\n", publicKey)
	fmt.Printf("Private Key: %s\n", privateKey)

	nodeClient, err := core.NewNodeClient(node)
	if err != nil {
		panic(err)
	}

	if nodeClient.CheckConnection() {

		data, err := nodeClient.GetWalletData(publicKey)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Balance: %d,%d", data.Balance.Integral, data.Balance.Fraction)

		nodeClient.CloseConnection()

		if nodeClient.CheckConnection() {
			panic(fmt.Errorf("connection is still open"))
		}

	} else {
		panic(fmt.Errorf("connection is not open"))
	}

}
