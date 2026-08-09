package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Реши квадратное уравнение")
	var a, b, c float64
	fmt.Println("Введи число a")
	fmt.Scan(&a)
	fmt.Println("Введи число b")
	fmt.Scan(&b)
	fmt.Println("Введи число c")
	fmt.Scan(&c)
	D := math.Pow(b, 2) - (4 * a * c)
	if D > 0 {
		var x1, x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b + math.Sqrt(D)) / (2 * a)
		fmt.Println("Уравнение имеет два корня\nD = " + fmt.Sprint(D))
		fmt.Println("X1 = " + fmt.Sprint(x1))
		fmt.Println("X1 = " + fmt.Sprint(x2))

	} else if D == 0 {
		var x float64
		x = -b / (2 * a)
		fmt.Println("X = " + fmt.Sprint(x))

	} else if D < 0 {
		fmt.Println("Решений нет\nD = " + fmt.Sprint(D))
	}
}
