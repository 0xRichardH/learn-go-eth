package common

import (
	"context"
	"regexp"

	"0x4richard.eth/learning/internal"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
)

func IsAddressValid(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func IsContract(address string) bool {
	ctx := context.Background()
	addr := ethCommon.HexToAddress(address)
	client, err := internal.NewLocalClient().Dial()
	if err != nil {
		zerolog.Ctx(ctx).Fatal().Err(err).Msg("failed to dial ethclient")
		panic(err)
	}

	byteCode, err := client.CodeAt(ctx, addr, nil)
	if err != nil {
		zerolog.Ctx(ctx).Fatal().Err(err).Msg("failed to get code at address")
		panic(err)
	}

	return len(byteCode) > 0
}
