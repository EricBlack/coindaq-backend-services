package wallet

import (
	"strconv"

	"github.com/btcsuite/btcutil"
)

func getBTCAmount(amount string) (btcutil.Amount, error) {
	famount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return btcutil.Amount(0), err
	}
	btcAmount, err := btcutil.NewAmount(famount)
	if err != nil {
		return btcutil.Amount(0), err
	}
	return btcAmount, nil
}
