package fruit

type FruitItem struct{
	Name string
	Weight float64
}

func NewFruitItem(name string, weight float64) FruitItem {
	return FruitItem{
		Name: name,
		Weight: weight,
	}
}

