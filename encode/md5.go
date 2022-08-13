package encode

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// Md5 生成md5+Hex字符串
// Deprecated: 使用新版的Md5Encode替代
func Md5(str string) string {
	ms := md5.Sum([]byte(str))
	return hex.EncodeToString(ms[:])
}

// Md5Encode 生成md5字符串
// Deprecated: 使用新版的Md5EC替代
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

type ECType string

var (
	Base64 ECType = "base64"
	Hex    ECType = "hex"
	Byte   ECType = ""
)

// Md5EC Md5后根据类型获取(hex | base64 | byte)三种结果
func Md5EC[S string | []byte](s S, ct ECType) S {
	var res string
	ms := md5.Sum([]byte(s))
	switch ct {
	case Base64:
		res = base64.StdEncoding.EncodeToString(ms[:])
	case Hex:
		res = hex.EncodeToString(ms[:])
	default:
		return S(ms[:])
	}
	return S(res)
}
