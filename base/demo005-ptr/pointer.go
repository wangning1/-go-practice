// 指针
package main

import "fmt"

func main() {
	house := "I have a house"
	//从指针获取获取指针指向的值
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("ptr address: %p\n", ptr)

	//对指针进行取值操作
	value := *ptr
	fmt.Printf("value type:%T\n", value)
	fmt.Printf("value : %s\n", value)

	//使用指针修改值
	// *”操作符作为右值时，意义是取指针的值；作为左值时，也就是放在赋值操作符的左边时，
	// 表示a指向的变量。其实归纳起来，“*”操作符的根本意义就是操作指针指向的变量。当操作在右值时，就是取指向变量的值；
	// 当操作在左值时，就是将值设置给指向的变量。
	*ptr = "this is a dog"
	value = *ptr
	fmt.Println(value)

	//创建指针的另一种方式
	// new()函数可以创建一个对应类型的指针，创建过程会分配内存。被创建的指针指向的值为默认值。
	str := new(string)
	*str = "hello go world"
	fmt.Printf("str type %T\n", str)
	fmt.Printf("str address %p\n", str)
	fmt.Printf("str value %s\n", *str)

}
