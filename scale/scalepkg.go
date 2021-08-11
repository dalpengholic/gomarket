package scale
import (
	"gomarket/biobag"
	"gomarket/fruit"
)

type Scale struct {
}

func NewScale() Scale {
	return Scale{}
}
// func (s Scale) ShowWeight(f fruit.FruitItem) float64{
// 	return f.Weight
// }


func (s Scale) ShowWeight(arg interface{}) float64{
	switch arg.(type) {
		
	case fruit.FruitItem:
		i := arg.(fruit.FruitItem)
		return i.Weight

	case biobag.BioBag:
		b := arg.(biobag.BioBag)
		totalWeight := 0.0
		for i :=0; i < b.Count;	i++ {
			totalWeight += b.List[i].Weight
		}		
		return totalWeight
	default:
		return 0.0

	}

}