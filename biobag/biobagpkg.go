package biobag

import (
	"gomarket/fruit"
)

type BioBag struct {
	Count int
	Bag [5]*fruit.FruitItem
}

func NewBioBag() BioBag {
	return BioBag{
		Count: 0,
	}
}