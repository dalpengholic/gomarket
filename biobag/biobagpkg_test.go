package biobag_test
import (
	"testing"
	"gomarket/biobag"
	"gomarket/fruit"
	"reflect"
	"fmt"
)

func TestNewBioBag(t *testing.T) {
	result := biobag.NewBioBag()
	result_type := reflect.TypeOf(result).String()
	fmt.Println(result_type)
	if result_type != "biobag.BioBag" {
		t.Errorf("타입은 Scale 이어야 합니다.")
	}
}

func TestAdd(t *testing.T){
	biobag1 := biobag.NewBioBag()
	fruit1 := fruit.NewFruitItem("apple", 0.5)
	fruit2 := fruit.NewFruitItem("apple", 0.4)
	fruit3 := fruit.NewFruitItem("apple", 0.6)
	fruit4 := fruit.NewFruitItem("apple", 0.5)
	fruit5 := fruit.NewFruitItem("apple", 0.5)
	fruit6 := fruit.NewFruitItem("banana", 0.5)
	result1 := biobag1.Add(&fruit1)
	result2 := biobag1.Add(&fruit2)
	result3 := biobag1.Add(&fruit3)
	result4 := biobag1.Add(&fruit4)
	result5 := biobag1.Add(&fruit5)
	result6 := biobag1.Add(&fruit6)

	if result1 && result2 && result3 && result4 && result5 != true {
		t.Errorf("결과는 true 이어야 합니다.")
	}

	if result6 == true {
		t.Errorf("결과는 False 여야합니다.")
	}

}