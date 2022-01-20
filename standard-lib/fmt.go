//标准包 fmt <参考：https://www.liwenzhou.com/posts/Go/go_fmt/>
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//向外输出
	// 1.Print系<Print系列函数会将内容输出到系统的标准输出，
	// 区别在于Print函数直接输出内容，
	// Printf函数支持格式化输出字符串，
	// Println函数会在输出内容的结尾添加一个换行符。>
	fmt.Print("哈哈哈")
	fmt.Printf("那啥%s\n", "过来吧")
	fmt.Println("你说啥呢")

	// 2.Fprint<Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，
	//          我们通常用这个函数往文件中写入内容>
	fmt.Fprintln(os.Stdout, "向标准输出列入内容")
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}

	fmt.Fprintln(fileObj, "向文件中写入内容")

	// 3.Sprint <Sprint系列函数会把传入的数据生成并返回一个字符串>
	name := "小张"
	age := 20
	ret := fmt.Sprintf("%s's age is %v", name, age)
	fmt.Println("ret ", ret)
	// 4.Errorf <Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。>
	err2 := fmt.Errorf("这是一个错误")
	fmt.Println("err2 ", err2)

	e := errors.New("标准错误")
	w := fmt.Errorf("wrap error %w", e)
	fmt.Println(w)

	//格式化占位符
	// *printf系列函数都支持format格式化参数
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	//获取输入
	var (
		firstName string
		lastName  string
	)

	fmt.Scan(&firstName, &lastName)
	fmt.Printf("name %s,%s", firstName, lastName)

}
