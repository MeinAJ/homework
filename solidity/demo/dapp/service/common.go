package service

import "os"

var (
	SepoliaURL   = os.Getenv("SEPOLIA_URL")
	PrivateKey   = os.Getenv("WALLET_PRIVATE_KEY")
	ContractAddr = os.Getenv("CONTRACT_ADDRESS")
)
