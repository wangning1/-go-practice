// 并发同步-互斥锁
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count      int
	countGuard sync.Mutex
)

func CountAdd() {
	countGuard.Lock()
	count++
	countGuard.Unlock()
}

func GetCount() int {
	countGuard.Lock()
	defer countGuard.Unlock()
	return count
}

func main() {
	for i := 1; i < 20; i++ {
		go func() {
			CountAdd()
			fmt.Println(GetCount())
		}()
	}

	time.Sleep(time.Second * 5)
	fmt.Println("main主线程结束count=", GetCount())
}
