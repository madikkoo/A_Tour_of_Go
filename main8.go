package main

import (
	"fmt"
	"os"
) //есеп: сандар косындысы, сандардын ен улкени жане ен кишиси

func main() {
	var a, b, c, d, f, g int

	fmt.Print("Введите число a: ")
	fmt.Fscan(os.Stdin, &a)

	fmt.Print("Введите число b: ")
	fmt.Fscan(os.Stdin, &b)

	fmt.Print("Введите число c: ")
	fmt.Fscan(os.Stdin, &c)

	fmt.Print("Введите число d: ")
	fmt.Fscan(os.Stdin, &d)

	fmt.Print("Введите число f: ")
	fmt.Fscan(os.Stdin, &f)

	fmt.Print("Введите число g: ")
	fmt.Fscan(os.Stdin, &g)

	fmt.Println(a, b, c, d, f, g)

	result := sum(a, b, c, d, f, g)
	fmt.Println("Сумма чисел:", result)

	max := max_min(a, b, c, d, f, g)
	fmt.Println("Максимальное число:", max)

	min := min_min(a, b, c, d, f, g)
	fmt.Println("Минимальное число:", min)
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func max_min(nums ...int) int {
	max := nums[0] // Initialize max with the first number

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func min_min(nums ...int) int {
	min := nums[0] // Initialize min with the first number

	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}
