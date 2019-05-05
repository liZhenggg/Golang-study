package main

import "fmt"

//测试GO函数
func main() {
	arr1 := []int{1, 2, 3}
	fmt.Println("sum:", slice_sum(arr1))

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println("sum:", slice_sum2(arr2))

	var arr3 = []int{3, 2, 3, 1, 6, 4, 8, 9}
	fmt.Println(slice_sum3(arr3))

	fmt.Println(sum(100, 1))
	fmt.Println(sum(100, 1, 2))
	var arr4 = []int{1, 2, 3, 4, 5}
	fmt.Println(sum(100, arr4...))

	closureFunc()
	closureFunc2()
}

// 一般函数
func slice_sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 返回值两边的括号在只声明一个返回值类型的时候可以省略
// 可以为返回值预先定义一个名称，在函数结束的时候，直接使用return就可以返回所有的预定义返回值。
func slice_sum2(arr []int) (sum int) {
	sum = 0
	for _, v := range arr {
		sum += v
	}
	return
}

// 函数多返回值
func slice_sum3(arr []int) (sum int, avg float64) {
	sum = 0
	avg = 0
	for _, v := range arr {
		sum += v
	}
	avg = float64(sum) / float64(len(arr))
	return
}

// 可变长参数列表
func sum(base int, arr ...int) int {
	sum := base
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 闭包函数
func closureFunc() {
	var arr1 = []int{1, 2, 3, 4, 5}
	var base = 300
	var sum = func(arr ...int) int {
		total_sum := base
		for _, v := range arr1 {
			total_sum += v
		}
		return total_sum
	}
	fmt.Println("sum:", sum(arr1...))
}

// 闭包函数2
func closureFunc2() {
	var base = 0
	inc := func() {
		base += 1
	}
	fmt.Println("base:", base)
	inc()
	fmt.Println("base:", base)
}

// 闭包函数3
func closureFunc3() {

}
