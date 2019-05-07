package main

import (
	"fmt"

	math "net.duokr/math"
)

func main() {
	var a = 100
	var b = 200
	fmt.Println("Add demo:", math.Add(a, b))
	fmt.Println("Substract demo:", math.Subtract(a, b))
	fmt.Println("Multiply demo:", math.Multiply(a, b))
	fmt.Println("Divide demo:", math.Divide(a, b))
}
