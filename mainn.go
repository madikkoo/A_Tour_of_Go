package main

import (
	"fmt"
)

func UltimatePoinOne(n ***int) {
	**n = 1
}
func main() {
	a := 0
	fmt.Println(&a, a)
	b := &a
	fmt.Println(&b, b)
	n := &b
	fmt.Println(&n, n)
	UltimatePoinOne(&n)
	fmt.Println(a)
}
