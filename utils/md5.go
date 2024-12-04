package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// md5小写加密
// 注:查下crypto/md5的用法即可，这里知道有这个包就好了
func EncodeMd5(value string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(value))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// 转大写
func EncodeMD5(value string) string {
	return strings.ToUpper(EncodeMd5(value))
}

// md5加盐
func CreatePass(str, salt string) string {
	s := str + salt
	return EncodeMD5(s)
}

// 验证加盐密码
func ValidPassword(str, salt, password string) bool {
	ValidPassword := CreatePass(str, salt)
	// fmt.Println("ValidPassword:", ValidPassword)
	// fmt.Println("pass:", password)
	return ValidPassword == password
}
