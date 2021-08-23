package scale_test

import (
	"testing"
	"gomarket/scale"
	"reflect"
	"fmt"
	"gomarket/fruit"
	"gomarket/biobag"
)

func TestNewScale(t *testing.T){
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
	testItem3.Add(&testItem1)
	testItem3.Add(&testItem2)
	result1 := scale.ShowWeight(testItem1)
	result2 := scale.ShowWeight(testItem3)
	if result1 != 0.95 {
		t.Errorf("결과는 0.95이어야 합니다.")
	}
	if result2 != 1.47 {
		t.Errorf("결과는 1.47이어야 합니다.")
	}

}