package scale
import (
	"gomarket/biobag"
	"gomarket/fruit"
	"gomarket/tag"
	"time"
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

func (s Scale) PrintTag(arg interface{}) tag.ItemTag{
	tag1 := tag.NewTag()
	dt := time.Now()
	tag1.CurrentTime = dt.Format("2006-Jan-02 15:04:05 Monday")
	switch arg.(type) {
		
	case fruit.FruitItem:
		i := arg.(fruit.FruitItem)
		tag1.Name = i.Name
		tag1.TotalWeight = i.Weight
		return tag1

	case biobag.BioBag:
		b := arg.(biobag.BioBag)
		totalWeight := 0.0
		for i :=0; i < b.Count;	i++ {
			totalWeight += b.List[i].Weight
		}	
		tag1.Name = b.List[0].Name
		tag1.TotalWeight = totalWeight	
		return tag1

	default:
		return tag1

	}


}
