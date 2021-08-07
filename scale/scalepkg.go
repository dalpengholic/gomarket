package scale
import "gopractice/test2_210807/fruit"

type Scale struct {
}

func NewScale() Scale {
	return Scale{}
}
func (s Scale) ShowWeight(f fruit.FruitItem) float64{
	return f.Weight
}