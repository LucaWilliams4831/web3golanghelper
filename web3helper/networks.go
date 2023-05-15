package web3helper

type EVMNetwork struct {
	HttpUrl      string
	WebsocketUrl string
	ChainID      uint64
}

var AvalancheMainnet = &EVMNetwork{
	HttpUrl:      "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/avalanche/mainnet",
	WebsocketUrl: "wss://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/avalanche/mainnet/ws",
	ChainID:      43114,
}

var AvalancheFujiTesnet = &EVMNetwork{
	HttpUrl:      "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/avalanche/testnet",
	WebsocketUrl: "wss://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/avalanche/testnet/ws",
	ChainID:      43113,
}

var BinanceSmartChainMainnet = &EVMNetwork{
	HttpUrl:      "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/mainnet",
	WebsocketUrl: "wss://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/mainnet/ws",
	ChainID:      56,
}

var BinanceSmartChainTestnet = &EVMNetwork{
	HttpUrl:      "https://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/testnet",
	WebsocketUrl: "wss://speedy-nodes-nyc.moralis.io/84a2745d907034e6d388f8d6/bsc/testnet/ws",
	ChainID:      97,
}
