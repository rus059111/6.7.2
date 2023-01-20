package main

//Задание 2. Сумма двух чисел по единице

import (
	"fmt"
	//"math"
)

func main() {

	var number1 int
	var number2 int

	fmt.Println("Введите число 1")
	fmt.Scan(&number1)
	fmt.Println("Введите число 2")
	fmt.Scan(&number2)

	j := number1
	c := number1 + number2

	for {
		j = j + 1
		if j == c {
			fmt.Println(j)
			fmt.Println("test git22222")
			break
		}
	}
}
