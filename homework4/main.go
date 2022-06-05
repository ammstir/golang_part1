package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertSort(numbers []int) []int {
	if len(numbers) > 1 {
		for idx := 1; idx < len(numbers); idx++ {
			key := numbers[idx]
			jdx := idx - 1
			for ; jdx >= 0 && numbers[jdx] > key; jdx-- {
				numbers[jdx+1] = numbers[jdx]

			}
			numbers[jdx+1] = key
		}
	}

	return numbers
}

func main() {
	var numbers []int

	fmt.Print("Введите числа через пробел: ")
	in := bufio.NewReader(os.Stdin)
	rawNumbers, _ := in.ReadString('\n')

	splitedNumbers := strings.Fields(rawNumbers)
	fmt.Println(splitedNumbers)

	for _, strNumber := range splitedNumbers {
		number, err := strconv.Atoi(strNumber)
		if err == nil {
			numbers = append(numbers, number)
		}

	}

	numbers = insertSort(numbers)

	fmt.Println(numbers)

}
