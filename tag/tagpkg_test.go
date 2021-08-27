package tag_test

import (
	"fmt"
	"gomarket/tag"
	"reflect"
	"testing"
)

func TestNewTag(t *testing.T) {
	result := tag.NewTag()
	result_type := reflect.TypeOf(result).String()
	fmt.Println(result_type)
	if result_type != "tag.ItemTag" {
		t.Errorf("타입은 ItemTag 이어야 합니다.")
	}
}
