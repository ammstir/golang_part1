package main

import (
	"fmt"
	"math"
)

func main() {
	var s, d, l float64

	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&s)

	d = math.Sqrt(s/math.Pi) * 2
	fmt.Printf("Диаметр круга с площадью %.2f равен %.2f\n", s, d)

	l = math.Pi * d
	fmt.Printf("Длина окружности круга с площадью %.2f равна %.2f\n", s, l)
}
