package tag

import (
	// "fmt"
	// "gomarket/fruit"
	// "gomarket/biobag"
)

type ItemTag struct {
	Name string
	CurrentTime string
	TotalWeight float64
	PricePerWeight float64
	TotalPrice float64
}

func NewTag() ItemTag {
	return ItemTag{}
}




