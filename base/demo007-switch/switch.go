// switch
package main

import (
	"fmt"
	"time"
)

func getResult(args ...int) bool {
	for _, vaule := range args {
		if vaule < 60 {
			return false
		}
	}
	return true
}

func main() {
	//基本的使用
	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}
	//case 接多个值
	now := time.Now()
	switch now.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("周末")
	default:
		fmt.Println("工作日")
	}

	//switch 可以不接表达式，相当于if-else
	switch {
	case now.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	//Switch 可以接一个函数
	math := 80
	english := 100
	chinese := 99

	switch getResult(math, english, chinese) {
	case true:
		fmt.Println("所有成绩及格")
	default:
		fmt.Println("存在不及格门课")
	}

	//switch的穿透能力<fallthrough 只能穿透一层>
	s := "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}

}
