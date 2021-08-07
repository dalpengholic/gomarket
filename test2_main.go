package main
import (
    "fmt"
	"gopractice/test2_210807/fruit"
	"gopractice/test2_210807/scale"
	"unsafe"
)

func main(){
	fruit1 := fruit.NewFruitItem("banana", 0.3)
	fmt.Println(fruit1)
	fmt.Println(unsafe.Sizeof(fruit1))
	scale1 := scale.NewScale()
	weight := scale1.ShowWeight(fruit1)
	fmt.Println(weight)
	
}