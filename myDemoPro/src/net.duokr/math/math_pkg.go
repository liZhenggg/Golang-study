package math

//Go规定了只有首字母大写的函数才能从包导出使用，即其他调用这个包中函数的代码只能调用那些导出的函数。

func Add(a, b int) int {
	return a + b
}
func Subtract(a, b int) int {
	return a - b
}
func Multiply(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	if b == 0 {
		panic("Can not divided by zero")
	}
	return a / b
}
