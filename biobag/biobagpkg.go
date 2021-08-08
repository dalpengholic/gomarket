package biobag

import (
	"gomarket/fruit"
	"fmt"
)

// Biobag for containing all fruit items
type BioBag struct {
	Count int
	List [5]*fruit.FruitItem
}

func (b *BioBag) Add(item *fruit.FruitItem) bool {
	if b.Count < 4 {
		b.List[b.Count] = item
		b.Count += 1
		return true
	} else {
		fmt.Println("Bag이 다 찼습니다. 새로운 Bag에 넣어주세요.")
		return false
	}
}

func (b *BioBag) Remove() bool {
	if b.Count == 0 {
		fmt.Println("Bag이 비어있습니다. 비울수 없습니다.")
		return false
	} else {
		b.List[b.Count-1] = nil
		b.Count -= 1
		return true
	}
}

// Biobag factory
func NewBioBag() BioBag {
	return BioBag{
		Count: 0,
	}
}