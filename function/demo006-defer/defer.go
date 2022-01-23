// defer 延迟函数
package main

import (
	"fmt"
	"os"
)

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return 0
	}

	size := info.Size()
	return size

}

func main() {
	//defer 函数执行时机是当前函数将执行defer函数，按照defer函数的逆序进行执行
	// 即先被defer的语句最后执行，最后defer的函数，最先被执行

	fmt.Println("defer begin")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("defer end")

	// 使用延迟函数可以在函数退出时释放资源
	size := fileSize("./test.txt")
	fmt.Println("文件的大小：", size)

}
