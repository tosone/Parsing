package convertPinyin

import (
	"testing"
)

func TestConvert(t *testing.T) {
	res := Convert("世界你好")
	realRes := []string{"shi", "jie", "ni", "hao"}
	for k, v := range res {
		if v != realRes[k] {
			t.Errorf("convert(世界你好) failed. Got %v, expected %v.", res, realRes)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert("世界你好")
	}
}
