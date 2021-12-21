package main

import (
	"fmt"
	"github.com/smartcontractkit/substrate-adapter/adapter"
)

func main() {
	fmt.Println("Starting Substrate adapter")

	//privkey := os.Getenv("SA_PRIVATE_KEY")
	//txType := os.Getenv("SA_TX_TYPE")
	//endpoint := os.Getenv("SA_ENDPOINT")
	//port := os.Getenv("SA_PORT")

	privkey := "0x02f7e19423a36e4e0a0e095e97e7edd2dc8d051ae9da00c457d524e54713b07b"
	txType := "immortal"
	endpoint := "ws://54.209.240.158:8844"
	port := "8081"

	adapterClient, err := adapter.NewSubstrateAdapter(privkey, txType, endpoint)
	if err != nil {
		fmt.Println("Failed starting Substrate adapter:", err)
		return
	}

	adapter.RunWebserver(adapterClient.Handle, port)
}
