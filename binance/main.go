package binance

import "github.com/iwinston/goex/v2/binance/spot"

type Binance struct {
	Spot *spot.Spot
}

func New() *Binance {
	return &Binance{
		Spot: spot.New(),
	}
}
