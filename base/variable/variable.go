// 变量
// 五种变量声明方法
package main

import "fmt"

func main() {
	// 第一种 一行声明一个变量
	var name string
	fmt.Println("name = ", name)
	name = "wang"
	fmt.Println("name = ", name)

	// 第二种 多个变量一起声明
	var (
		addr    string
		company string
	)

	addr = "北京"
	company = "jd"

	fmt.Println("addr = ", addr)
	fmt.Println("company = ", company)

	// 第三种 声明和初始化一个变量
	age := 12
	fmt.Println("age = ", age)

	// 第四种 声明和初始化多个变量
	school, subject := "清华", "物理"
	fmt.Println("school = ", school)
	fmt.Println("subject = ", subject)

	// 第五种 new函数声明一个指针变量
	// 使用表达式 new(Type) 将创建一个Type类型的匿名变量，初始化为Type类型的零值，然后返回变量地址，返回的指针类型为*Type。
	// &后面接变量名，表示取出该变量的内存地址
	// * 后面接指针变量，表示从内存地址中取出值
	ptr := new(int)
	fmt.Println("ptr address = ", ptr)
	fmt.Println("ptr value = ", *ptr)
	*ptr = 20
	fmt.Println("ptr value = ", *ptr)

}
