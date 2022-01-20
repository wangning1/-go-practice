// 函数的声明
package main

import "fmt"

type InnerData struct {
	a int
}

type Data struct {
	complex  []int
	ptr      *InnerData
	instance InnerData
}

//普通函数声明
func person(age int, name string) string {
	info := fmt.Sprintf("%s's age is %d", name, age)
	return info
}

//函数声明参数类型简写
func add(a, b int) int {
	return a + b
}

//函数多返回值
func multiRes(a int, b string) (int, string) {
	fmt.Println("inner a address", &a)
	return a, b
}

//Go语言中传入和返回参数在调用和返回时都使用值传递，这里需要注意的是指针、
// 切片和map等引用型对象指向的内容在参数传递中不会发生复制，而是将指针进行复制，类似于创建一次引用。
func passByValue(param Data) Data {
	//输出参数成员情况
	fmt.Printf("param value: %+v\n", param)
	//输出参数指针
	fmt.Printf("param ptr: %p\n", &param)
	return param
}

func main() {

	a := 2
	b := 3
	fmt.Println("inner a address", &a)
	result := add(a, b)
	fmt.Println("add res = ", result)

	name := "小王"
	age := 13
	info := person(age, name)
	fmt.Println("info = ", info)

	score, subject := multiRes(100, "math")
	fmt.Printf("subject %s's score is %d\n", subject, score)

	//测试passByValue函数
	in := Data{
		complex: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	//输出结构体成员情况
	fmt.Printf("in value %+v\n", in)
	//输出结构体指针
	fmt.Printf("in ptr: %p\n", &in)

	out := passByValue(in)
	//输出结果成员情况
	fmt.Printf("out value %+v\n", out)
	//输出结果指针
	fmt.Printf("out ptr %p\n", &out)

}
