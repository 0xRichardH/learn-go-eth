package internal

import "github.com/ethereum/go-ethereum/ethclient"

type ClientInterface interface {
	Dial() (*ethclient.Client, error)
}
