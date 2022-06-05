package main

import "fmt"

func main() {
	var number int
	var hund, tens, ones string

	fmt.Print("Введите трёхзначное число: ")
	fmt.Scanln(&number)

	if number < 100 || number > 999 {
		fmt.Println("Вы ввели не трёхзначное число.")
	} else {
		d1 := number % 10
		d2 := (number / 10) % 10
		d3 := number / 100

		if d1 == 1 {
			ones = "единица"
		} else if d1 < 5 && d1 > 1 {
			ones = "единицы"
		} else {
			ones = "единиц"
		}

		if d2 == 1 {
			tens = "десяток"
		} else if d2 < 5 && d2 > 1 {
			tens = "десятка"
		} else {
			tens = "десятков"
		}

		if d3 == 1 {
			hund = "сотня"
		} else if d3 < 5 && d3 > 1 {
			hund = "сотни"
		} else {
			hund = "сотен"
		}

		fmt.Printf("В введённом числе %d %s, %d %s и %d %s\n", d3, hund, d2, tens, d1, ones)
	}
}
