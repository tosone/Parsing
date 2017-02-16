package betterMatch

import (
	"fmt"
	"testing"
)

func TestBetterMatch(t *testing.T) {
	res1 := Match("想加", []string{"想家"})
	if res1 != "想家" {
		t.Errorf("Match(想加). Except: %v, Real: %v", "想家", res1)
	}
	fmt.Println(MatchWhole("不好", []string{"什么", "不好"}))
	res2 := MatchWhole("想加", []string{"下架", "想家"})
	if checkSlice(res2, []string{"想家", "下架"}) {
		t.Errorf("Match(想加). Except: %v, Real: %v", "想家", res2)
	}
}

func checkSlice(str1, str2 []string) bool {
	if len(str1) == len(str2) {
		return false
	}
	for k, v := range str1 {
		if str2[k] != v {
			return true
		}
	}
	return false
}
