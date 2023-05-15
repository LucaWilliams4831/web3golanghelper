package main

import (
	"context"
	"fmt"

	"github.com/nikola43/web3golanghelper/web3helper"
)

func main() {

	rpcUrl := "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/testnet"
	web3HttpClient := web3helper.NewHttpWeb3Client(rpcUrl)

	chainID, err := web3HttpClient.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Chain Id: " + chainID.String())
}
