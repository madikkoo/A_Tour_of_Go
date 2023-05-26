package main

import (
	"fmt"
)

func main() {
	l := Strlen("Hello world!")
	fmt.Println(l)
}
func Strlen(s string) int {
	return len(s)
}
