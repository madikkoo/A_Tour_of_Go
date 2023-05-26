package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Sam Harris"
	initials := GetInitials(name)
	fmt.Println(initials)
}

func GetInitials(name string) string {
	words := strings.Split(name, " ") // Split the name into words
	initials := ""
	for _, word := range words {
		initials += strings.ToUpper(string(word[0])) + "."
	}
	initials = strings.TrimSuffix(initials, ".") // Remove the trailing dot
	return initials
}
