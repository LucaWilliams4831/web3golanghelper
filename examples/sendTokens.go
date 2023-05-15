package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/nikola43/web3golanghelper/web3helper"
)

func main2() {

	rpcUrl := "https://eth-goerli.nodereal.io/v1/703500179cfc4348b90bebc0b3fba854"
	wsUrl := "wss://eth-goerli.nodereal.io/ws/v1/703500179cfc4348b90bebc0b3fba854"
	pk := "cfedfad8629f43cfffda1bc9a4c97e1aa4461615f8331b0760272f9303b2838e"
	web3Helper := web3helper.NewWeb3GolangHelper(rpcUrl, wsUrl)
	

	chainID, err := web3Helper.HttpClient().ChainID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	//

	tx, nonce, err := web3Helper.SendTokens("0xc43aF0698bd618097e5DD933a04F4e4a5A806834", "0x6AD058b6af6BEEF79a20174D9f651f3534Fe2F60", big.NewInt(1000000000000000000), pk)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("tx", tx)
	fmt.Println("nonce", nonce)

	fmt.Println("Chain Id: " + chainID.String())
}
