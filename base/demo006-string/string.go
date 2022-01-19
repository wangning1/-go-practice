//字符串应用
package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	str := "hello,欢迎!"
	//字符串的长度, len()函数的返回值的类型为int，表示字符串的ASCII字符个数或字节长度。
	//Go语言的字符串都以UTF-8格式保存，每个中文占用3个字节，因此使用len()获得两个中文文字对应的6个字节
	fmt.Println("str byte len is:", len(str))

	//统计字符个数
	fmt.Printf("str char len is:%d", utf8.RuneCountInString(str))

	//遍历字符串 unicode字符串遍历使用for range
	for _, s := range str {
		fmt.Printf("unicode %c %d\n", s, s)
	}

	//获取在字符串中的下标
	co := strings.Index(str, ",")
	fmt.Println(co)

	//获取字符串中的某一段字符
	fmt.Println(str[co:])

	//修改字符串,字符串一旦声明是不可变的，可以转为字节数组，然后修改字节数组的方式进行修改字符串
	//byte和string可以进行强制类型转换
	angle := "hero never die"
	angleByte := []byte(angle)
	for i := 5; i <= 10; i++ {
		angleByte[i] = ' '
	}
	fmt.Println("angleByte to str:", string(angleByte))

	//字符串链接的两种方式
	//1.+
	name := "王" + "宁"
	fmt.Println("name:", name)
	//2.字节缓存
	var strBuf bytes.Buffer
	strBuf.WriteString("hello ")
	strBuf.WriteString("wang ning")
	fmt.Println("strBuf", strBuf.String())

}
