package funcDemo1

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

//函数声明
func f1(input1 string, input2 int32) (output1 string, output2 uint32) { return "", 11 }
func f2(string, int32) (result string)                                { result = ""; return }
func f3(string) string                                                { return "1" + "2" }

//函数类型声明
type MyFunc func(string, int32) string

//MyFunc的实现
func myFunc(str string, in int32) string { return "aa" }

// 等价于 var splice MyFunc
var splice func(string, int32) string

// splice = myFunc
// splice("abc",123)

////
//员工Id生成器
type EmployeeGenerator func(company string, department string, sn uint32) string

// 默认公司名称
var company = "Gophers"

// 序列号
var sn uint32

//生成员工Id
func generatoId(generator EmployeeGenerator, department string) (string, bool) {
	// 若员工ID生成器不可用，则无法生成员工ID，应直接返回。
	if generator == nil {
		return "", false
	}

	// 使用代码包 sync/atomic 中提供的原子操作函数可以保证并发安全。
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，所以提供这样一个函数作为辅助。
func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

func FunDemo1() {
	var generator EmployeeGenerator
	generator = func(company string, department string, sn uint32) string {
		return appendSn(company+"-"+department+"-", sn)
	}
	//Gophers-RD-1 true。
	fmt.Println(generatoId(generator, "RD"))
}
