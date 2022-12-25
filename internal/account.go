package internal

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Account struct {
	client  ethclient.Client
	address string
}

var _ AccountInterface = (*Account)(nil)

func NewAccount(client ethclient.Client, address string) AccountInterface {
	return &Account{client: client, address: address}
}

func (a *Account) Address() common.Address {
	return common.HexToAddress(a.address)
}

func (a *Account) Balance(ctx context.Context) (*big.Float, error) {
	balance, err := a.client.BalanceAt(ctx, a.Address(), nil)
	if err != nil {
		return nil, fmt.Errorf("account: failed to get balance: %w", err)
	}

	return convertBalanceToDecimal(balance), nil
}

func (a *Account) PendingBalance(ctx context.Context) (*big.Float, error) {
	balance, err := a.client.PendingBalanceAt(ctx, a.Address())
	if err != nil {
		return nil, fmt.Errorf("account: failed to get pending balance: %w", err)
	}

	return convertBalanceToDecimal(balance), nil
}

func convertBalanceToDecimal(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	if _, isSuccess := fbalance.SetString(balance.String()); isSuccess == false {
		return nil
	}

	ethBalance := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethBalance
}
