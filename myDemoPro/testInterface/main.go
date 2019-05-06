package main

import "fmt"

// 接口测试类
func main() {
	//一个类型A只要实现了接口X所定义的全部方法，那么A类型的变量也是X类型的变量。
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(ApplePhone)
	phone.call()

	var phones = [5]Phone{
		NokiaPhone{1000},
		NokiaPhone{2000},
		ApplePhone{3000},
		ApplePhone{4000},
		ApplePhone{5000},
	}

	totalPrice := 0
	for _, p := range phones {
		totalPrice += p.sale()
	}
	fmt.Println("totalPrice:", totalPrice)

	//接口类型还可以作为结构体的数据成员
	var person = Person{append(phones[:], ApplePhone{6000}), "Jams", 28}
	fmt.Println("person.name:", person.name)
	fmt.Println("person.age:", person.age)
	fmt.Println("person.totalCost:", person.totalCost())
}

type Phone interface {
	call()
	sale() int
}

type NokiaPhone struct {
	price int
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I'm Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) sale() int {
	return nokiaPhone.price
}

type ApplePhone struct {
	price int
}

func (applePhone ApplePhone) call() {
	fmt.Println("I'm Apple, I can call you!")
}

func (applePhone ApplePhone) sale() int {
	return applePhone.price
}

type Person struct {
	phones []Phone
	name   string
	age    int
}

func (person Person) totalCost() (sum int) {
	sum = 0
	for _, v := range person.phones {
		sum += v.sale()
	}
	return
}
