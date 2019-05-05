// myDemoPro project main.go
package main

import (
	"fmt"
	//funcs "myDemoPro/funcDemos"
	"strconv"
	"time"
)

var v_name = "Lion"
var v_name2 = v_name + "2"

//int string
func demo1() {
	// var v_name = "Lion2"
	var hobby = "codeing"
	var hobby2 = hobby + "And sleep"
	hobby = "aaa"

	var age1 = 10
	age2 := age1
	age1 += 1

	// var book1 String = "“历史的真相”"
	book2 := "“历史的真相2”"

	fmt.Println(v_name2+" like "+hobby2+",age:", age2, "!", &age1, &age2, "I reading books:"+book2)

}

//int string 格式化
func demo2() {
	var num1 uint64 = 22233
	var num2 int32 = -22332222
	size := (8)

	fmt.Println("num1:", num1, "num2:", num2)
	fmt.Printf("类型为 uint64 的整数 %d 需占用的存储空间为：%d 个字节。\n", num1, size)

	var num3 int = -0x1000
	fmt.Printf("16进制数 %X 表示的是 %d 。\n", num3, -4096)

	var num4 = 5.89E-4
	fmt.Printf("浮点数 %E 表示的是 %f。 \n", num4, 0.000589)

	fmt.Printf("99用浮点数float32表示为: %f。 \n", float32(99))

	var num5 = 3.7E+1 + 5.98E-2i
	fmt.Printf("浮点数 %E 表示的是 %f。 \n", num5, (37.000000 + 0.059800i))

	var char1 rune = '赞'
	fmt.Printf("字符 '%c' 的Unicode代码是 %s。\n", char1, ("U+8D5E"))
}

// string 格式化
func demo3() {
	var str1 string = "\\\""
	fmt.Printf("用解释型字符串表示法表示的 %q 所代表的是 %s。\n", str1, `\"`)
}

//数组
func demo4() {
	var arr [5]int
	arr[0] = 6
	arr[3] = arr[0] - 3
	arr[1] = arr[2] + 5
	arr[4] = len(arr)

	sum := 19

	fmt.Printf("%v \n", (sum == arr[0]+arr[1]+arr[2]+arr[3]+arr[4]))
}

//切片(slice)
func demo5() {
	// []int{1, 2, 3} //声明

	var arr = [...]int{9, 5, -14, 108, 3, -2, 13, 45}
	slice1 := arr[1:len(arr)]           // [5, -14, 108, 3, -2, 13, 45]
	slice2 := slice1[2 : len(slice1)-2] // [108, 3, -2]
	length := (3)
	capacity := (5)
	fmt.Printf("%v, %v\n", (length == len(slice2)), (capacity == cap(slice2)))

	fmt.Printf("cap(slice2): %v \n", cap(slice2))

	// runtime error: index out of range
	// var str string = strconv.Itoa(slice2[0]) + "," + strconv.Itoa(slice2[1]) +
	// 	"," + strconv.Itoa(slice2[2]) + "," + strconv.Itoa(slice2[3])

	// runtime error: slice bounds out of range
	// slice3 := slice1[2 : len(slice1)-2 : 4]

	slice3 := slice1[2 : len(slice1)-2 : 5]
	// runtime error: index out of range
	// var str string = strconv.Itoa(slice3[0]) + "," + strconv.Itoa(slice3[1]) +
	// 	"," + strconv.Itoa(slice3[2]) + "," + strconv.Itoa(slice3[3])

	slice3 = append(slice3, 33, 44, 55)
	var str string = strconv.Itoa(slice3[0]) + "," + strconv.Itoa(slice3[1]) +
		"," + strconv.Itoa(slice3[2]) + "," + strconv.Itoa(slice3[3]) + "," + strconv.Itoa(slice3[4]) +
		"," + strconv.Itoa(slice3[5])

	fmt.Println("slice2: " + str)

	var slice4 = []int{0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Print("befor copy: slice4:")
	for i := 0; i < len(slice4); i++ {
		fmt.Print(strconv.Itoa(slice4[i]) + ",")
	}
	copy(slice4, slice3)
	fmt.Print("\nafter copy: slice4:")
	for i := 0; i < len(slice4); i++ {
		fmt.Print(strconv.Itoa(slice4[i]) + ",")
	}

}

//切片操作
func demo6() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := numbers4[4:6:8]
	length := (2)
	capacity := (4)
	fmt.Printf("%v, %v\n", length == len(slice5), capacity == cap(slice5))
	slice5 = slice5[:cap(slice5)]
	slice5 = append(slice5, 11, 12, 13)
	length = (7)
	fmt.Printf("%v\n", length == len(slice5))
	slice6 := []int{0, 0, 0}
	copy(slice5, slice6)
	e2 := (0)
	e3 := (8)
	e4 := (11)
	fmt.Printf("%v, %v, %v\n", e2 == slice5[2], e3 == slice5[3], e4 == slice5[4])
}

//字典类型 Map<k> = v
func demo7() {
	//声明 赋值
	// map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	mm := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	b := mm[3]
	fmt.Println("mm[3] value:", b)

	mm[3] = b + "3"
	fmt.Println("mm[3] newvalue:", mm[3])

	//修改添加
	mm[4] = "ddd"
	mm[5] = "eee"
	fmt.Println("mm[4] value:", mm[4])
	fmt.Println("mm[5] value:", mm[5])
	fmt.Println("mm[6] value:", mm[6])

	//空值
	e, ok := mm[6]
	fmt.Printf("mm[6] value: %v ,是否有值: %v \n", e, ok)

	//删除
	delete(mm, 5)
	delete(mm, 6)
	e, ok = mm[5]
	fmt.Printf("mm[5] value: %v ,是否有值: %v \n", e, ok)
}

//缓存通道
func demo8() {
	// make声明
	// sli1 := make([]string{"a", "b", "c"}, 3)
	// map1 := make(map[int]string{1: "ff1", 2: "ff2"})

	ch1 := make(chan string, 3)

	ch1 <- "value1"
	fmt.Printf("ch1.length= %v \n", len(ch1))

	ch1 <- "value2"
	ch1 <- "value3"
	fmt.Printf("ch1.length= %v \n", len(ch1))

	v1, ok := <-ch1
	fmt.Printf("ch1 value: %v ,是否有值: %v \n", v1, ok)

	close(ch1)
	// close(ch1)
	v1, ok = <-ch1
	fmt.Printf("ch1 value: %v ,是否有值: %v \n", v1, ok)

	ch2 := make(chan string, 3)

	go func() {
		ch2 <- ("已到达！")
	}()
	var value string = "数据"
	value = value + <-ch2

	fmt.Printf("value: %v \n", value)
}

//非缓冲通道
func demo9() {
	// chan3 := make(chan string, 0)

	////单向通道
	//接受通道
	type Receiver <-chan int
	//发送通道
	type Sender chan<- int

	var myChannel = make(chan int, 2)
	var sender Sender = myChannel
	var receiver Receiver = myChannel
	sender <- 11
	fmt.Printf("sender length: %v \n", len(sender))

	//error:receive from send-only type Sender
	// var val1 = <-sender

	//error:send to receive-only type Receiver
	// receiver <- 22

	var val2 = <-receiver
	fmt.Printf("receiver val: %v \n", val2)

	//test demo
	var myChannel2 = make(chan int, 0)
	var number = 6
	go func() {
		var sender Sender = myChannel2
		sender <- number
		fmt.Println("Sent!")
	}()
	go func() {
		var receiver Receiver = myChannel2
		fmt.Println("Receiver!", <-receiver)
	}()
	// 让main函数执行结束的时间延迟1秒，
	// 以使上面两个代码块有机会被执行。
	time.Sleep(time.Second)
}

func main() {
	// fmt.Println("Hello World!")
	// demo1()
	// demo2()
	// demo3()
	// demo4()
	// demo5()
	// demo6()
	// demo7()
	demo8()
	// demo9()
	// funcs.FunDemo1()
}
