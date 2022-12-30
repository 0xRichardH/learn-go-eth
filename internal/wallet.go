package internal

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog"
)

type Wallet struct {
	PrivateKEY *ecdsa.PrivateKey
	PublicKEY  *ecdsa.PublicKey
}

var _ WalletInterface = (*Wallet)(nil)

func NewWallet(ctx context.Context) WalletInterface {
	w := &Wallet{}
	if err := w.generateKeys(); err != nil {
		zerolog.Ctx(ctx).Fatal().Err(err).Msg("failed to generate keys")
		return nil
	}

	return w
}

func (w *Wallet) GenerateNewAddress() string {
	return crypto.PubkeyToAddress(*w.PublicKEY).Hex()
}

func (w *Wallet) generateKeys() error {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return fmt.Errorf("failed to generate private key: %w", err)
	}

	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("failed to get public key")
	}

	w.PrivateKEY = privateKey
	w.PublicKEY = publicKey

	return nil
}
