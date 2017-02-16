package timeline

import "testing"

func TestConvert(t *testing.T) {
	res1 := Convert("200")
	if res1 != 200 {
		t.Errorf("Convert(200) failed. Got %d, expected 200.", res1)
	}
	res2 := Convert("20'0")
	if res2 != 1200 {
		t.Errorf("Convert(20'0) failed. Got %d, expected 1200.", res2)
	}
	res3 := Convert("0'80")
	if res3 != 60 {
		t.Errorf("Convert(0'80) failed. Got %d, expected 60.", res3)
	}
	res4 := Convert("1'20")
	if res4 != 80 {
		t.Errorf("Convert(0'80) failed. Got %d, expected 80.", res4)
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Convert(string(i) + "'" + string(i))
	}
}
