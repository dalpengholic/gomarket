package scale_test

import (
	"testing"
	"gomarket/scale"
	"reflect"
	"fmt"
)

func TestNewScale(t *testing.T){
	result := scale.NewScale()
	result_type := reflect.TypeOf(result).String()
	fmt.Println(result_type)
	if result_type != "scale.Scale" {
		t.Errorf("타입은 Scale 이어야 합니다.")
	}

}