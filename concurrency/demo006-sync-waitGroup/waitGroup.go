// 同步-线程间等待
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.baidu.com",
		"http://www.jd.com",
		"https://www.wanmei.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			content, err := http.Get(url)
			fmt.Printf("url:%s;content:%v;err %s\n", url, content, err)
		}(url)
	}

	wg.Wait()
	fmt.Println("request over!")
}
