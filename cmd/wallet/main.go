package main

import (
	"context"
	"fmt"

	"0x4richard.eth/learning/internal"
)

func main() {
	ctx := context.Background()
	wallet := internal.NewWallet(ctx)
	fmt.Println(wallet.GenerateNewAddress())
}
