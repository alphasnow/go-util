package encode

import (
	"reflect"
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

func TestMd5Encode(t *testing.T) {
	// admin 21232f297a57a5a743894a0e4a801fc3
	actual := Md5Encode("admin", "base64")
	expect := "ISMvKXpXpadDiUoOSoAfww=="
	if actual != expect {
		t.Error(actual, expect)
	}
}

func TestMd5EC(t *testing.T) {

	t.Run("encode string", func(t *testing.T) {
		actual := Md5EC[string]("admin", Hex)
		expect := "21232f297a57a5a743894a0e4a801fc3"
		if actual != expect {
			t.Error(actual, expect)
		}
	})

	t.Run("encode byte", func(t *testing.T) {
		actual := Md5EC("admin", Hex)
		expect := "21232f297a57a5a743894a0e4a801fc3"
		if actual != expect {
			t.Error(actual, expect)
		}
	})

	t.Run("encode byte", func(t *testing.T) {
		actual := Md5EC([]byte("admin"), Byte)
		kind := reflect.TypeOf(actual).Kind()
		if kind != reflect.Slice {
			t.Error(kind, reflect.Slice)
		}
	})
}
