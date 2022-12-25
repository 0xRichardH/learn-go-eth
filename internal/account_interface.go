package internal

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type AccountInterface interface {
	Address() common.Address
	Balance(ctx context.Context) (*big.Float, error)
	PendingBalance(ctx context.Context) (*big.Float, error)
}
