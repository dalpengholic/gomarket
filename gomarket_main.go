package main

import (
	"database/sql"
	"fmt"
	"gomarket/biobag"
	"gomarket/fruit"
	"gomarket/platform/fruitdb"
	"gomarket/scale"
	"gomarket/tag"
	"unsafe"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create items
	fruit1 := fruit.NewFruitItem("banana", 0.3)
	fruit2 := fruit.NewFruitItem("apple", 0.28)
	fruit3 := fruit.NewFruitItem("Orange", 0.9)
	fmt.Println("생성된 과일 인스턴스1:", fruit1)
	fmt.Println(unsafe.Sizeof(fruit1))

	// Create a scale
	scale1 := scale.NewScale()

	// Check weight of the item
	weight1 := scale1.ShowWeight(fruit1)
	fmt.Println("scale1에 측정한 과일인스턴스1의 무게: ", weight1)
	weight2 := scale1.ShowWeight(fruit2)
	fmt.Println("scale1에서 측정한 과일인스턴스2의 무게: ", weight2)

	tag1 := tag.NewTag()
	fmt.Println("This is a new tag instance: ", tag1)

	// Create a biobag
	biobag1 := biobag.NewBioBag()
	fmt.Println("새로운 biobag 인스턴스:", biobag1)

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
	fmt.Println("result5: ", result5)
	// Test PrintTag with biobag
	result6 := scale1.PrintTag(biobag1)
	fmt.Println("result6: ", result6)

	db, _ := sql.Open("sqlite3", "./fruit.db")
	defer db.Close()
	fruitDB := fruitdb.NewFruitInfo(db)

	fruitDB.Add(fruitdb.Item{
		Name:  "banana",
		Price: 0.4,
	})

	fruitDB.Add(fruitdb.Item{
		Name:  "tomato",
		Price: 0.4,
	})

	items := fruitDB.Get()
	fmt.Println(items)

}
