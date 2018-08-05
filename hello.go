package main

import (
	"fmt"
	"github.com/golang-demos/go-concat-strings-library/strlib"
)

func main() {
	var finalStr = strlib.ConcatStrs("Hello, ", "Go!")
	fmt.Println(finalStr)
}
