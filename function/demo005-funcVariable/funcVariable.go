// 函数的可变参数
package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer

	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		var typeString string

		switch s.(type) {
		case bool:
			typeString = "bool"
		case int:
			typeString = "int"
		case string:
			typeString = "string"
		}

		b.WriteString("value:")
		b.WriteString(str)

		b.WriteString(" type:")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	fmt.Println(printTypeValue("wang3", 12, true))
}
