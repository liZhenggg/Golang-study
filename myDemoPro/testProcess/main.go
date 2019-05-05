package main

import "fmt"

//GO中的流程控制语句
func main() {
	// testIfElse()

	// testSwitch()

	// testFor()

}

// test for: if..else if..else
func testIfElse() {
	var dog_age = 2
	var dog_sex = "M"

	if dog_age < 0 {
		fmt.Println("is a baby dog")
	} else if dog_age >= 1 && dog_age <= 10 {
		fmt.Println("is a small dog")
	} else {
		fmt.Println("is a big dog")
	}

	if dog_sex == "M" && dog_age > 1 {
		fmt.Println("is a M dog and age greater than 1")
	}
}

//test for: switch
func testSwitch() {
	// score 为 [0,100]之间的整数
	var score int = 69
	switch score / 10 {
	case 10:
		fmt.Println("完美")
	case 9:
		fmt.Println("优秀")
	case 8:
		fmt.Println("良好")
	case 7:
		fmt.Println("一般")
	case 6:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

}

// test for: for
func testFor() {
	var i int = 1
	for ; i <= 100; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	//死循环
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }
}
