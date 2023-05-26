package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Running tests...")

	testAbbrevName("Sam Harris", "S.H")
	testAbbrevName("Patrick Feenan", "P.F")
	testAbbrevName("Evan Cole", "E.C")
	testAbbrevName("P Favuzzi", "P.F")
	testAbbrevName("David Mendieta", "D.M")

	fmt.Println("Tests complete.")
}

func testAbbrevName(name string, expected string) {
	result := AbbrevName(name)
	if result == expected {
		fmt.Printf("PASS: AbbrevName(%q) = %q\n", name, result)
	} else {
		fmt.Printf("FAIL: AbbrevName(%q) = %q, expected %q\n", name, result, expected)
	}
}

func AbbrevName(name string) string {
	words := strings.Split(name, " ")
	initials := ""
	for _, word := range words {
		initials += strings.ToUpper(string(word[0])) + "."
	}
	initials = strings.TrimSuffix(initials, ".")
	return initials
}
