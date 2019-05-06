package main

import "fmt"

func main() {
	// //计算矩形结构体的面积
	// var rect Rect2
	// rect.width = 3
	// rect.length = 5
	// fmt.Println("rect area:", rect.width*rect.length)

	// //使用初始化方式为结构体赋值
	// var rect2 = Rect2{width: 4, length: 5}
	// fmt.Println("rect2 area:", rect2.width*rect2.length)

	// //按照结构体成员定义的顺序赋值
	// var rect3 = Rect2{5, 5}
	// fmt.Println("rect3 area:", rect3.width*rect3.length)

	// //结构体参数也是值传递
	// fmt.Println("double_area(rect3):", double_area(rect3))
	// fmt.Println("rect3 area:", rect3.width*rect3.length)

	// //使用结构体组合函数(结构体的内部方法)--method
	// fmt.Println("rect3.area_mothod():", rect3.area_mothod())

	// var rect3_2 = new(Rect2)
	// rect3_2.width = 5
	// rect3_2.length = 5
	// fmt.Println("rect3_2.area_mothod():", rect3_2.area_mothod())

	// //结构体和指针
	// var rect4 = new(Rect2)
	// rect4.width = 5
	// rect4.length = 6
	// fmt.Println("rect4.area():", rect4.area())
	// fmt.Println("width:", rect4.width, ",length:", rect4.length)

	//结构体内嵌类型
	var p IPhone
	p.phone.price = 6000
	p.phone.color = "Balck"
	p.model = "IPhone 8P"
	fmt.Println("price:", p.phone.price)
	fmt.Println("color:", p.phone.color)
	fmt.Println("model:", p.model)

	//可以直接使用内部结构体属性
	var p2 Samsung
	p2.price = 5000
	p2.color = "white"
	p2.model = "S9"
	fmt.Println("price:", p2.price)
	fmt.Println("color:", p2.color)
	fmt.Println("model:", p2.model)
}

//定义一个矩形结构体
type Rect struct {
	width  float64
	length float64
}

// 简化
type Rect2 struct {
	width, length float64
}

//go函数的参数是值传递，结构体参数也是值传递
func double_area(rect Rect2) float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

//结构体组合函数
func (rect Rect2) area_mothod() float64 {
	return rect.width * rect.length
}

//结构体和指针，主要作用是在函数内部改变传递进来变量的值
func (rect *Rect2) area() float64 {
	rect.width *= 2
	rect.length *= 2
	return rect.width * rect.length
}

type Phone struct {
	price int
	color string
}

type IPhone struct {
	phone Phone
	model string
}

//可以不再定义变量，直接引入结构体
type Samsung struct {
	Phone
	model string
}
