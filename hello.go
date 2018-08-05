package main

import (
	"fmt"
	"bytes"
)

func concatStrs(strs ...string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(strs); i++ {
		buffer.WriteString(strs[i])
	}
	return buffer.String()
}

func main() {
	var finalStr = concatStrs("Hello, ", "Go!")
	fmt.Println(finalStr)
}
