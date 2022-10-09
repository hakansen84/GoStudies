package main

import (
	"fmt"
	"strings"
)

func main() {
	newString := "Go is a great programming language. Go for it!"

	fmt.Println(strings.HasPrefix(newString, "Go"))

}
