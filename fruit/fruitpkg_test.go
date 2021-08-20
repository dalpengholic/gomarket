package fruit_test
import "testing"
import "gomarket/fruit"

func TestNewFruitItem(t *testing.T){
	result1 := fruit.NewFruitItem("apple", 0.7)
	if result1.Name != "apple" {
		t.Errorf("결과1 값은 apple 이어야 합니다.")
	}
	if result1.Weight != 0.7 {
		t.Errorf("결과1 무게는 0.7 이어야 합니다.")
	}
	

	result2 := fruit.NewFruitItem("banana", 0.7)
	if result2.Name != "banana" {
		t.Errorf("결과1 값은 banana 이어야 합니다.")
	}
	if result2.Weight != 0.7 {
		t.Errorf("결과2 값은 0.7 이어야 합니다")
	}
}
