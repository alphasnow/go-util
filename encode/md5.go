package encode

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// Md5 生成md5+Hex字符串
// Deprecated: 使用新版的Md5Encode替代
func Md5(str string) string {
	return Md5Encode(str, "hex")
}

// Md5Encode 生成md5字符串
func Md5Encode(str string, typ string) string {
	var res string
	ms := md5.Sum([]byte(str))
	switch typ {
	case "base64":
		res = base64.StdEncoding.EncodeToString(ms[:])
	default:
		res = hex.EncodeToString(ms[:])
	}
	return res
}
