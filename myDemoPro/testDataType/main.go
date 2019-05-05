package main

import (
	"fmt"
)

func main() {
	// fmt.Println("123")
	// testArray()
	// testSlice()
	// appendFunc()
	testMap()

}

// 数组 Array
func testArray() {
	//// 定义：
	// 1.先定义后赋值
	var x [3]int
	x[0] = 11
	x[1] = 12
	x[2] = 13
	fmt.Println("x:", x)

	// range 函数可以用来遍历数组，切片和字典
	for i, v := range x {
		fmt.Printf(">%v:%v\n", i, v)
	}

	// 2.定义同时赋初始值
	var x2 = [3]int{21, 22, 23}
	fmt.Println("x2:", x2)

	// range 函数当未使用到索引时，应当使用下划线代替(_).
	for _, v := range x2 {
		fmt.Println(">>", v)
	}

	// 3.先定义变量后赋值
	var x3 = [3]int{}
	x3[0] = 31
	x3[1] = 32
	x3[2] = 33
	fmt.Println("x3:", x3)

	//如果将数组元素定义在不同行上面，那么最后一个元素后面必须跟上}或者,。
	var x4 = [...]string{"Monday", "Tuesday",
		"Wednesday",
	}
	fmt.Println("x4:", x4)

	var x5 = [...]string{"Monday", "Tuesday",
		"Wednesday"}
	fmt.Println("x5:", x5)
}

// 切片 Slice
func testSlice() {
	//// 定义
	// 1. 先声明一个变量是切片，然后使用内置函数make去初始化这个切片

	//只指定长度，这个时候切片的长度和容量是相同的
	var x = make([]float64, 5)
	fmt.Println("Capcity:", cap(x), "Length:", len(x))

	var y = make([]float64, 5, 10)
	fmt.Println("Capcity:", cap(y), "Length:", len(y))

	// 2.通过取数组切片来赋值:[low_index:high_index]
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var s1 = arr1[2:3]
	var s2 = arr1[:3]
	var s3 = arr1[2:]
	var s4 = arr1[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	var arr2 = make([]int, 5, 10)
	for i := 0; i < len(arr2); i++ {
		arr2[i] = i
	}
	fmt.Println("arr2:", arr2)
	fmt.Println("Capcity:", cap(arr2), ",Length:", len(arr2))

	// 内置函数append
	arr2 = append(arr2, 5, 6, 7, 8)
	fmt.Println("arr2:", arr2)
	fmt.Println("Capcity:", cap(arr2), ",Length:", len(arr2))

	//Go在默认的情况下，如果追加的元素超过了容量大小，Go会自动地重新为切片分配容量，容量大小为原来的两倍。
	arr2 = append(arr2, 9, 10, 11, 12)
	fmt.Println("arr2:", arr2)
	fmt.Println("Capcity:", cap(arr2), ",Length:", len(arr2))

}

// append 函数
func appendFunc() {
	var slice1 = append([]int{1, 2, 3}, 4, 5, 6)
	fmt.Println("slice1:", slice1)

	slice2 := append([]int{1, 2, 3}, []int{4, 5, 6}...)
	fmt.Println("slice2:", slice2)

	bytes := append([]byte("hello"), "world"...)
	fmt.Println("bytes:", bytes)
}

// 字典 Map
func testMap() {
	// 定义方式：
	// 1.初始化数据
	var map1 = map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Cherry",
		"D": "Durian",
	}

	for key, val := range map1 {
		fmt.Println("Key:", key, ",Value:", val)
	}

	// 2.使用make函数
	var map2 map[string]string
	map2 = make(map[string]string)
	map2["A"] = "Apple"
	map2["B"] = "Banana"
	map2["C"] = "Cherry"
	map2["D"] = "Durian"

	for _, val := range map2 {
		fmt.Println("Value:", val)
	}

	// 使用make定义简化写法
	map3 := make(map[string]string)

	// 访问不存在的键将返回零值
	map4 := make(map[string]int)
	map4["A"] = 0
	map4["B"] = 12
	map4["C"] = 28

	fmt.Println("map3[\"A\"]:", map3["A"])
	fmt.Println("map4[\"A\"]:", map4["A"])
	fmt.Println("map4[\"D\"]:", map4["D"])

	// 区分字典中的零值和真实值
	// map[Key]的返回值有2个：值和是否存在此键的bool型变量。可根据bool型变量是否为true来判断键是否存在
	if val, ok := map4["C"]; ok {
		fmt.Println("map4[\"C\"] value:", val)
	}
	if val, ok := map4["D"]; ok {
		fmt.Println("map4[\"D\"] value:", val)
	}

	fmt.Println("Before Delete：")
	fmt.Println("Length:", len(map4))
	fmt.Println(map4)
	delete(map4, "A")
	fmt.Println("After Delete：")
	fmt.Println("Length:", len(map4))
	fmt.Println(map4)

	// 复杂字典结构
	var students = make(map[string]map[string]int)
	students["0001"] = map[string]int{"Andy": 25}
	students["0002"] = map[string]int{"Bill": 22}
	students["0003"] = map[string]int{"Carl": 23}
	for stu_no, stu_info := range students {
		fmt.Println("stu_no:", stu_no)
		for stu_name, stu_age := range stu_info {
			fmt.Println("stu_name:", stu_name, ",stu_age:", stu_age)
		}
	}

	students2 := map[string]map[string]int{
		"20001": {"Andy": 25},
		"20002": {"Bill": 22},
		"20003": {"Carl": 23},
	}

	for stu_no, stu_info := range students2 {
		fmt.Println("stu_no:", stu_no)
		for stu_name, stu_age := range stu_info {
			fmt.Println("stu_name:", stu_name, ",stu_age:", stu_age)
		}
	}
}
