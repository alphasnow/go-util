package encode

import (
	"fmt"
)

func ExampleMd5() {
	res := Md5("admin")
	fmt.Println(res)
	// Output: 21232f297a57a5a743894a0e4a801fc3
}

func ExampleMd5Encode() {
	res := Md5Encode("admin", "base64")
	fmt.Println(res)
	// Output: ISMvKXpXpadDiUoOSoAfww==
}

func ExampleMd5EC() {
	res := Md5EC("admin", "base64")
	fmt.Println(res)
	// Output: ISMvKXpXpadDiUoOSoAfww==
}
