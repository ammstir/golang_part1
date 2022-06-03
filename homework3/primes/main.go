package main

import (
	"fmt"
)

func update_primes_with_number(number int, primes []int) []int {
	for _, prime := range primes {
		if prime < 2 {
			continue
		}
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

	fmt.Print("Введите положительное число: ")
	fmt.Scanln(&number)

	if number < 0 {
		fmt.Println("Это не положительное число!")
		return
	}

	for i := 0; i <= number; i++ {
		primes = update_primes_with_number(i, primes)
	}

	fmt.Println(primes)
}
