package blockchain

import (
	"fmt"
	"log"

	"go-carbon-backend/contracts" // 由 abigen 生成的包

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	infuraURL    = "https://sepolia.infura.io/v3/22c489e1df7a458181cc8ddc54e70f1b"
	ContractAddr = "0x71f9cbd442304a86ba905968fea9d250d9f08b04"
)

// BlockchainConnector 封装了与以太坊的连接和合约实例
type BlockchainConnector struct {
	Client   *ethclient.Client
	Contract *contracts.Contracts
}

// NewConnector 创建一个新的区块链连接器实例
func NewConnector() (*BlockchainConnector, error) {
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	address := common.HexToAddress(ContractAddr)
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate contract: %v", err)
	}

	log.Println("Successfully connected to Sepolia and instantiated contract.")

	return &BlockchainConnector{
		Client:   client,
		Contract: instance,
	}, nil
}
