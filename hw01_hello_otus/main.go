package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	hello := "Hello, OTUS!"

	reversedString := stringutil.Reverse(hello)

	fmt.Println(reversedString)
}
