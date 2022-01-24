// 使用接口进行数据排序
package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"3.test",
		"2.hello",
		"4.golang",
		"1.what",
		"5.how",
	}

	sort.Sort(names)
	for _, val := range names {
		fmt.Println("val:", val)
	}
}
