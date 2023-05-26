package main

import (
	"fmt"
)

func main() {
	s := "Hello world!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}