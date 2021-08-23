package biobag_test
import (
	"testing"
	"gomarket/biobag"
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