package main

import (
	"fmt"
)

func update_primes_with_number(number int, primes []int) []int {
	for _, prime := range primes {
		if prime > (number / 2) {
			break
		}
		if number%prime == 0 {
			return primes
		}
	}

	primes = append(primes, number)

	return primes
}

func main() {
	var number int
	var primes []int

	fmt.Print("Введите положительное число больше 1: ")
	fmt.Scanln(&number)

	if number <= 1 {
		fmt.Println("Это не положительное число или оно не больше 1!")
		return
	}

	for i := 2; i <= number; i++ {
		primes = update_primes_with_number(i, primes)
	}

	fmt.Println(primes)
}
