package huobi

import (
	"github.com/iwinston/goex/v2/huobi/futures"
	"github.com/iwinston/goex/v2/huobi/spot"
)

type HuoBi struct {
	Spot    *spot.Spot
	Futures *futures.Futures
}

func New() *HuoBi {
	return &HuoBi{
		Spot:    spot.New(),
		Futures: futures.New(),
	}
}
