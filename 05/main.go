package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "We are happy."
	fmt.Println(replace(input))
}

func replace(input string) string {
	return strings.Replace(input, " ", "%20", -1)
}
