package main

import (
	"fmt"
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

/*
This function will divide the int a and b.
The result of this division will be stored in the int pointed by div.
The remainder of this division will be stored in the int pointed by mod.
*/
func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // Perform division and store the result in the variable pointed by div
	*mod = a % b // Calculate the remainder and store it in the variable pointed by mod
}

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
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

func main() {
	n := 5
	fmt.Println(n)
}
