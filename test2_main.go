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
	fmt.Println(fruit1)
	fmt.Println(unsafe.Sizeof(fruit1))
	scale1 := scale.NewScale()
	weight := scale1.ShowWeight(fruit1)
	fmt.Println(weight)
	biobag1 := biobag.NewBioBag()
	fmt.Println(biobag1)
	
}