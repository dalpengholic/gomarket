package fruit
// import (
// 	"fmt"
// 	"unsafe"
// )


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

// func main() {
// 	fruit1 := NewFruitItem("banana", 0.3)
// 	fruit2 := NewFruitItem("apple", 0.4)
// 	fmt.Println(fruit1,fruit2)
// 	fmt.Printf("Address of fruit1: %p, Address of fruit2: %p\n", &fruit1, &fruit2)
// 	fmt.Println(unsafe.Sizeof(fruit1))
	// banana1 := NewItem("banana", 0.3)
	// banana2 := NewItem("banana", 0.6)
	// apple1 := NewItem("apple", 0.2)
	// var arr [3]*Item
	// arr[0], arr[1], arr[2] = &banana1, &banana2, &apple1

	// fmt.Println(arr)
	// fmt.Println()
	// fmt.Println(&banana1)

	// str := "ABC"
	// fmt.Println(&str)
// }