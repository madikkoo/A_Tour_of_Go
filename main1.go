package main

import "fmt"

func UltimatePointOne(n **int) {
	**n = 1
}

func main() {
	a := 0
	fmt.Println(&a, a)
	b := &a
	fmt.Println(&b, b)
	n := b
	fmt.Println(&n, n)
	UltimatePointOne(&n)
	fmt.Println(a)
}
