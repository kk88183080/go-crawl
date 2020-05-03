package main

import (
	"fmt"
	"regexp"
)

/**
正则表达式测试
*/
func main() {
	str := "i love 441553747@qq.com, 8818@163.com"
	// 完全匹配
	//compile := regexp.MustCompile("441553747@qq.com")
	// 只匹配以@qq.com结尾的邮箱
	compile := regexp.MustCompile(`[a-zA-Z0-9]*@qq.com`)

	s := compile.FindString(str)

	fmt.Println(s)
}
