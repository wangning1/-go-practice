// 初始化结构体
package main

import "fmt"

type People struct {
	name  string
	child *People
}

type Address struct {
	Provice  string
	City     string
	ZipCode  string
	PhoneNum string
}

func main() {
	// 初始化方式1：key-value键值对形式
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Printf("relation %+v\n", relation)

	//初始化方式2：使用多个值的列表初始化结构体
	add := &Address{
		"河南",
		"驻马店",
		"0396",
		"12321313",
	}
	fmt.Printf("add %+v\n", add)

	//匿名结构体
	msg := struct {
		id   int
		name string
	}{
		1024,
		"code",
	}
	fmt.Printf("msg %+v\n", msg)

}
