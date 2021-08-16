package main
import (
    "fmt"
	"gomarket/fruit"
	"gomarket/scale"
	"gomarket/biobag"
	"gomarket/tag"
	"unsafe"
)

func main(){
	// Create items
	fruit1 := fruit.NewFruitItem("banana", 0.3)
	fruit2 := fruit.NewFruitItem("apple", 0.28)
	fruit3 := fruit.NewFruitItem("Orange", 0.9)
	fmt.Println(fruit1)
	fmt.Println(unsafe.Sizeof(fruit1))

	// Create a scale
	scale1 := scale.NewScale()

	// Check weight of the item
	weight1 := scale1.ShowWeight(fruit1)
	fmt.Println(weight1)
	weight2 := scale1.ShowWeight(fruit2)
	fmt.Println(weight2)

	tag1 := tag.NewTag()
	fmt.Println("This is a new tag: ", tag1)

	// Create a biobag
	biobag1 := biobag.NewBioBag()
	fmt.Println(biobag1)

	// Add item into the bag
	result1 := biobag1.Add(&fruit1)
	fmt.Println(result1, biobag1)

	// Remove item into the bag
	// result2 := biobag1.Remove()
	// fmt.Println(result2, biobag1)
	result3 := biobag1.Add(&fruit2)
	result4 := biobag1.Add(&fruit3)

	fmt.Println(result3, biobag1, fruit2)
	biobag1.List[0].Name = "Orange"
	fmt.Println(*biobag1.List[0])

    fmt.Println(result4)
    fmt.Println("현재 비닐에 있는 아이템:", biobag1)
	fmt.Println("아이템들: ", biobag1.List[0], biobag1.List[1], biobag1.List[2])
	weight3 := scale1.ShowWeight(biobag1)
	fmt.Println(weight3)

    // test newtag
	tag2 := tag.NewTag()
	fmt.Println(tag2)

    // Test PrintTag with fruit item
	result5 := scale1.PrintTag(fruit1)
	fmt.Println("result5: " ,result5)
    // Test PrintTag with biobag 
	result6 := scale1.PrintTag(biobag1)
	fmt.Println("result6: " ,result6)



	

}