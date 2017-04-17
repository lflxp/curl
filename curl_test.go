package curl

import "testing"


func TestHttpGet(t *testing.T) {
	data := HttpGet("http://www.baidu.com")
	if len(data) != 0 {
		t.Log(data)
	} else {
		t.Fatal("nothing got")
	}
}

func BenchmarkHttpGet(b *testing.B) {
	for i:=0;i<b.N;i++ {
		data := HttpGet("http://www.baidu.com")
		if len(data) == 0 {
			b.Fatal("nothing got")
		} else {
			b.Log(data)
		}
	}
}
