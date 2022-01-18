// 流程控制if-else
package main

import "fmt"

func main() {
	//单分支判断

	age := 20
	gender := "male"
	if age > 20 {
		fmt.Println("已成年")
	}

	if age >= 20 && gender == "male" {
		fmt.Println("成年男性")
	}

	//多分支判断
	if age > 20 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}

	//高级写法<在 if 里可以允许先运行一个表达式，取得变量后，再对其进行判断，比如第一个例子里代码也可以写成这样>
	if score := 99; score > 90 {
		fmt.Println("成绩优秀")
	}

}
