package code

import (
	"crypto/md5"
	"fmt"
	"webshell/common"
)

func Md5(a string) string {
	b := md5.Sum([]byte(common.Password)) // 加密数据
	//fmt.Printf("%x",b) // 转换为16进制，并打印
	a = fmt.Sprintf("%x", b)
	return a
}
