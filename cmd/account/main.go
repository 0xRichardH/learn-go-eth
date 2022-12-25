package main

import (
	"context"
	"fmt"

	"0x4richard.eth/learning/internal"
	"github.com/rs/zerolog"
)

const (
	TestWalletAddress = "0xc044B36EC3Bc15dEfB719Ef233A39354c99E52b6"
)

func main() {
	ctx := context.Background()

	client, err := internal.NewLocalClient().Dial()
	if err != nil {
		zerolog.Ctx(ctx).Fatal().Err(err).Msg("failed to dial ethclient")
	}

	account := internal.NewAccount(*client, TestWalletAddress)
	address := account.Address()
	fmt.Println(address.Hex())
	fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())

	balance, errB := account.Balance(ctx)
	if errB != nil {
		zerolog.Ctx(ctx).Fatal().Err(errB).Msg("failed to get balance")
	}
	fmt.Println(balance)

	pendingBalance, errPB := account.PendingBalance(ctx)
	if err != nil {
		zerolog.Ctx(ctx).Fatal().Err(errPB).Msg("failed to get pending balance")
	}
	fmt.Println(pendingBalance)
}
