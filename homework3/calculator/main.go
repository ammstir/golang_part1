package main

import (
	"fmt"
	"math"
)

func main() {
	var operand1, operand2, result float64
	var operator string

	fmt.Print("Введите первый аргумент: ")
	fmt.Scanln(&operand1)
	fmt.Print("Введите математическое действие: ")
	fmt.Scanln(&operator)

	if operator != "!" {
		fmt.Print("Введите второй аргумент: ")
		fmt.Scanln(&operand2)
	}

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Нельзя делить на ноль!")
			return
		} else {
			result = operand1 / operand2
		}
	case "*":
		result = operand1 * operand2
	case "^":
		result = math.Pow(operand1, operand2)
	case "!":
		result = 1
		for i := 1; i <= int(operand1); i++ {
			result *= float64(i)
		}
	}

	fmt.Printf("Получилось %.2f\n", result)
}
