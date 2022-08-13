package encode

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	// admin 21232f297a57a5a743894a0e4a801fc3
	actual := Md5("admin")
	expect := "21232f297a57a5a743894a0e4a801fc3"
	if actual != expect {
		t.Error(actual, expect)
	}
}

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5("admin")
	}
}

func ExampleMd5() {
	res := Md5("admin")
	fmt.Println(res)
	// Output: 21232f297a57a5a743894a0e4a801fc3
}

func TestMd5Encode(t *testing.T) {
	// admin 21232f297a57a5a743894a0e4a801fc3
	actual := Md5Encode("admin", "base64")
	expect := "ISMvKXpXpadDiUoOSoAfww=="
	if actual != expect {
		t.Error(actual, expect)
	}
}

func BenchmarkMd5Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5Encode("admin", "base64")
	}
}

func ExampleMd5Encode() {
	res := Md5Encode("admin", "base64")
	fmt.Println(res)
	// Output: ISMvKXpXpadDiUoOSoAfww==
}
