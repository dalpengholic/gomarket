package main
import (
    "fmt"
	"gomarket/fruit"
	"gomarket/scale"
	"gomarket/biobag"
	"unsafe"
)

func main(){
	fruit1 := fruit.NewFruitItem("banana", 0.3)
	fruit2 := fruit.NewFruitItem("apple", 0.2)
	fmt.Println(fruit1)
	fmt.Println(unsafe.Sizeof(fruit1))
	scale1 := scale.NewScale()
	weight1 := scale1.ShowWeight(fruit1)
	fmt.Println(weight1)
	weight2 := scale1.ShowWeight(fruit2)
	fmt.Println(weight2)
	biobag1 := biobag.NewBioBag()
	fmt.Println(biobag1)
	result1 := biobag1.Add(&fruit1)
	fmt.Println(result1, biobag1)
	result2 := biobag1.Remove()
	fmt.Println(result2, biobag1)
	result3 := biobag1.Add(&fruit2)
	fmt.Println(result3, biobag1, fruit2)
	biobag1.List[0].Name = "Orange"
	fmt.Println(*biobag1.List[0])

}