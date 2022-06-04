package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insert_sort(numbers []int) []int {
	if len(numbers) > 2 {
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
	raw_numbers, _ := in.ReadString('\n')

	splited_numbers := strings.Fields(raw_numbers)
	fmt.Println(splited_numbers)

	for _, str_number := range splited_numbers {
		number, err := strconv.Atoi(str_number)
		if err == nil {
			numbers = append(numbers, number)
		}

	}

	numbers = insert_sort(numbers)

	fmt.Println(numbers)

}
