// go 轻量级线程
package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(1000)
	}
}

func main() {
	go running()

	var input string
	fmt.Scanln(&input)
}
