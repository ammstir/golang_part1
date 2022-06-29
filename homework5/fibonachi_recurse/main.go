package main

import "fmt"

func fibonachi(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return fibonachi(n-1) + fibonachi(n-2)
}

func main() {
	var number int

	fmt.Print("Введите положительное число: ")
	fmt.Scanln(&number)

	fmt.Printf(
		"Число фибоначи с номером %d - это %d\n",
		number,
		fibonachi(number),
	)
}
