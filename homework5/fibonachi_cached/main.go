package main

import "fmt"

var cache = make(map[int]int, 0)

func fibonachi(n int) int {
	if _, ok := cache[n]; ok {
		return cache[n]
	}

	switch n {
	case 1:
		return 0
	case 2:
		return 1
	default:
		cache[n-1] = fibonachi(n - 1)
		cache[n-2] = fibonachi(n - 2)

		return cache[n-1] + cache[n-2]
	}

}

func main() {
	var number int

	fmt.Print("Введите положительное число: ")
	fmt.Scanln(&number)

	if number < 0 {
		panic("Я не умею работать с отрицательными числами")
	}

	fmt.Printf(
		"Число фибоначи с номером %d - это %d\n",
		number,
		fibonachi(number),
	)
}
