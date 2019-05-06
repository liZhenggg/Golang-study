package main

import "fmt"

func main() {
	var x int

	var x_pointer *int

	x = 10
	x_pointer = &x

	fmt.Println("x：", x)
	fmt.Println("x_pointer：", x_pointer)
	fmt.Println("x的地址（x_pointer = &x）：", x_pointer)
	fmt.Println("通过指针变量获得x的值（*x_pointer）：", *x_pointer)
	fmt.Println("指针的地址（&x_pointer）：", &x_pointer)
	fmt.Println("*&x_pointer：", *&x_pointer)

}
