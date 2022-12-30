package internal

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type LocalClient struct {
	address   string
	ETHClient *ethclient.Client
}

var _ ClientInterface = (*LocalClient)(nil)

func NewLocalClient() ClientInterface {
	// return &LocalClient{address: "http://localhost:8545"}
	return &LocalClient{address: "http://localhost:8545"}
}

func (c *LocalClient) Dial() (*ethclient.Client, error) {
	client, err := ethclient.Dial(c.address)
	if err != nil {
		return nil, fmt.Errorf("ethclient: failed to dial: %w", err)
	}

	return client, nil
}
