package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Введите сторону прямоугольника: ")
	fmt.Scanln(&a)
	fmt.Print("Введите другую сторону прямоугольника: ")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугольника: %d\n", a*b)
}
