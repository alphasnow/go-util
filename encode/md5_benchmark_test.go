package encode

import (
	"testing"
)

func BenchmarkMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5("admin")
	}
}

func BenchmarkMd5Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5Encode("admin", "base64")
	}
}

func BenchmarkMd5Code(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Md5En([]byte("admin"), Hex)
	}
}
