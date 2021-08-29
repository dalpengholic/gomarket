package scale_test

import (
	"fmt"
	"gomarket/biobag"
	"gomarket/fruit"
	"gomarket/scale"
	"reflect"
	"testing"
)

func TestNewScale(t *testing.T) {
	result := scale.NewScale()
	result_type := reflect.TypeOf(result).String()
	fmt.Println(result_type)
	if result_type != "scale.Scale" {
		t.Errorf("타입은 Scale 이어야 합니다.")
	}

}

func TestShowWeight(t *testing.T) {
	scale := scale.NewScale()
	testItem1 := fruit.NewFruitItem("banana", 0.95)
	testItem2 := fruit.NewFruitItem("apple", 0.52)
	testItem3 := biobag.NewBioBag()
	testItem4 := biobag.NewBioBag()
	testItem3.Add(&testItem1)
	testItem3.Add(&testItem2)
	result1 := scale.ShowWeight(testItem1)
	result2 := scale.ShowWeight(testItem3)
	result3 := scale.ShowWeight(testItem4)
	if result1 != 0.95 {
		t.Errorf("결과는 0.95이어야 합니다.")
	}
	if result2 != 1.47 {
		t.Errorf("결과는 1.47이어야 합니다.")
	}
	if result3 != 0.0 {
		t.Errorf("결과는 0.0 이어야 합니다.")
	}
	result4 := scale.ShowWeight(100)
	if result4 != 0.0 {
		t.Errorf("결과는 0.0 이어야 합니다.")
	}

}

func TestPriceTag(t *testing.T) {
	scale := scale.NewScale()
	testItem1 := fruit.NewFruitItem("banana", 0.95)
	testItem2 := fruit.NewFruitItem("banana", 0.52)
	testItem3 := biobag.NewBioBag()
	testItem3.Add(&testItem1)
	testItem3.Add(&testItem2)
	result1 := scale.PrintTag(testItem1)
	if result1.Name != "banana" {
		t.Errorf("결과는 banana 이어야 합니다.")
	}
	if result1.TotalWeight != 0.95 {
		t.Errorf("결과는 0.95 이어야 합니다.")
	}
	result2 := scale.PrintTag(testItem3)
	if result2.Name != "banana" {
		t.Errorf("결과는 banana 이어야 합니다.")
	}
	if result2.TotalWeight != 1.47 {
		t.Errorf("결과는 1.47 이어야 합니다.")
	}

	result3 := scale.PrintTag(20)
	result3_type := reflect.TypeOf(result3).String()
	fmt.Println(result3_type)
	if result3_type != "tag.ItemTag" {
		t.Errorf("타입은 ItemTag 이어야 합니다.")
	}

}
