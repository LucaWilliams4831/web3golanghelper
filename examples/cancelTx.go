package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/nikola43/web3golanghelper/web3helper"
)

func main() {

	wsUrl := "wss://eth-goerli.nodereal.io/ws/v1/703500179cfc4348b90bebc0b3fba854"

	rpcUrl := "https://rpc2.sepolia.org"
	pk := "00a2bb2d32a9aa43994686051b4d9368a64ebc1df7f18ae85e7d043b486443b8"

	web3Helper := web3helper.NewWeb3GolangHelper(rpcUrl, wsUrl)
	web3Helper.AddAccount(pk)

	chainID, err := web3Helper.HttpClient().ChainID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	//

	tx, nonce, err := web3Helper.SendEth("0xf0d7f526c0a706745d9bbdbd66a1a36e0f40da0f", 1, pk)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(tx)


	tx, err = web3Helper.CancelTx("0xf0d7f526c0a706745d9bbdbd66a1a36e0f40da0f", big.NewInt(int64(nonce)), 10, pk)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("tx", tx)

	fmt.Println("Chain Id: " + chainID.String())
}
