package encode

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	ms := md5.Sum([]byte(str))
	return hex.EncodeToString(ms[:])
}
