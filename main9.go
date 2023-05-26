package main

import (
	"fmt"
	"math"
	"os"
)

// есеп: сан берилген сол санды тубир астына салу керек

func main() {
	var a int

	fmt.Println("Введите число:")
	_, err := fmt.Fscanf(os.Stdin, "%d", &a)
	if err != nil {
		fmt.Println("Ошибка при чтении числа:", err)
		return
	}
	sum := 0
	for i := 0; i <= a; i++ {
		fmt.Print(i, " ")
		sum += i
	}
	fmt.Println()
	fmt.Println("Сумма чисел:", sum)
	fmt.Println("Корень суммы:", math.Sqrt(float64(sum)))
}
