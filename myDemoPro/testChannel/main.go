package main

import (
	"fmt"
	"strconv"
	"time"
)

// 测试Channel
func main() {
	// Channel
	// testChannel()

	// 多个channel
	// testMultiChannel()

	// 带缓冲区的通道
	// testChannelBuffer()

	testChannelRang()

	var input string
	fmt.Scanln(&input)
}

// 定点跳投
func fixed_shooting(msg_chan chan string) {
	for {
		msg_chan <- "fixed shooting"
		// fmt.Println("fixed shooting！continue shooting>>")
	}
}

// 三分球
func three_point_shooting(msg_chan chan string) {
	for {
		msg_chan <- "3 point shooting"
		// fmt.Println("3 point shooting！continue shooting>>")
	}
}

// 统计
func count(msg_chan chan string) {
	for {
		msg := <-msg_chan
		fmt.Println("msg:", msg)
		time.Sleep(time.Second * 1)
	}
}

func testChannel() {
	var c chan string
	c = make(chan string)

	go fixed_shooting(c)
	go three_point_shooting(c)
	//要向channel里面写信息，必须有配对的取信息的一端，否则是不会写的。
	go count(c)
}

func testMultiChannel() {
	c_fixed := make(chan string)
	c_3_point := make(chan string)

	go fixed_shooting(c_fixed)
	go three_point_shooting(c_3_point)

	go func() {
		fmt.Println("begin......")
		for i := 0; i < 50; i++ {
			// for {
			fmt.Println(">>> i =", i)
			select {
			case msg1 := <-c_fixed:
				fmt.Println("fixed shooting!", msg1)
			case msg2 := <-c_3_point:
				fmt.Println("3 point shooting!", msg2)
			default:
				fmt.Println("select default>>>")
			}
		}
		fmt.Println("end......")
	}()

}

func shooting(msg_chan chan string) {
	for group := 1; group <= 5; group++ {
		fmt.Println("------- group:", group)
		for i := 1; i <= 10; i++ {
			msg_chan <- strconv.Itoa(group) + ":" + strconv.Itoa(i)
		}
		time.Sleep(time.Second * 3)
	}
}

func getChannel(msg_chan chan string) {
	for {
		fmt.Println(<-msg_chan)
	}
}

func testChannelBuffer() {
	var c = make(chan string, 20)
	go shooting(c)
	go getChannel(c)
}

func testChannelRang() {
	chan1 := make(chan string, 2)
	chan1 <- "hello"
	chan1 <- "world"
	close(chan1)

	for ele := range chan1 {
		fmt.Println(ele)
	}

	val, ok := <-chan1
	fmt.Println("val1:", val, ",ok:", ok)
}
