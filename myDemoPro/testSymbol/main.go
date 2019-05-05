package main

import "fmt"

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) size() float64 {
	return r.width * r.height
}

//Go中的*和&符号
func main() {
	fmt.Println("main")

	// & 是取地址符号，&Rect是取到Rect类型对象的地址
	fmt.Println(&Rect{100, 100})

	// * 是指针运算符，可以表示变量是指针类型，也可以表示指针类型变量所指向的存储单元，也就是这个地址所指向的值
	var r *Rect = &Rect{100, 100}
	fmt.Println(r)

	fmt.Println(*r)

	// 查看这个指针变量的地址，基本数据类型直接打印地址
	fmt.Println(&r)

	// type WebSite struct {
	// 	Name string
	// }
	// ws1 := WebSite{"CC"}
	// ws2 := WebSite{"Brown"}

	// var websites = []WebSite{ws1, ws2}
	// var ws *WebSite
	// for _, v := range websites {
	// 	if v.Name == "Brown" {
	// 		ws = &v
	// 	}
	// }
	// fmt.Println(ws.Name)

	// fmt.Println(&Rect{100, 100})

}
