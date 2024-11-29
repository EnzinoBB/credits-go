package test

import (
	"fmt"
	"testing"

	"github.com/EnzinoBB/credits-go/core"
	"github.com/EnzinoBB/credits-go/model"
	"github.com/EnzinoBB/credits-go/utils"
)

func TestGetWalletData(t *testing.T) {

	//A Credits Node API End-Point
	node := "localhost:9090"

	//generate new Credits wallet using utility feature
	publicKey, privateKey, err := utils.GenerateKeys()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Printf("Public Key: %s\n", publicKey)
	fmt.Printf("Private Key: %s\n", privateKey)

	nodeClient, err := core.NewNodeClient(node)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if nodeClient.CheckConnection() {

		data, err := nodeClient.GetWalletData(publicKey)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		data_out := model.GetWalletData_Out(data)

		fmt.Printf("Wallet Data:\n %v", data_out)

		nodeClient.CloseConnection()

		if nodeClient.CheckConnection() {
			t.Errorf("connection is still open")
			t.Fail()
		}

	} else {
		t.Errorf("connection is not open")
		t.Fail()
	}

}
