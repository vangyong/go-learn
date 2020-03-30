package main

import (
	. "grammar"
)

func main() {

	//引用类型和值类型
	//PrintVariableDifferent()
	//fmt.Println("PrintVariableDifferent end")

	//结构
	//	var book1 Book
	//	book1.author = ""
	//	book1.author = "www.runoob.com"
	//	book1.subject = "Go 语言教程"
	//	book1.bookId = 6495407
	//	printStruct(book1)

	//数组
	// PrintArray()

	//循环遍历计算
	//	PrintRangeSum()

	//切片
	//	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//	var sum = sumSlice(numbers)

	//接口
	//pc := PhoneConnector{"PhoneConnecter"}
	//var a Connector
	//a = Connector()
	//a.Connect()

	//并发-goroutine
	go GoroutineSay("world")
		GoroutineSay("hello")

	//并发-channel
	//	s := []int{7, 2, 8, -9, 4, 0}
	//	c := make(chan int)
	//	go ChannelSum(s[:len(s)/2], c)
	//	go ChannelSum(s[len(s)/2:], c)
	//	x, y := <-c, <-c // 从通道 c 中接收
	//	fmt.Println(x, y, x+y)

	//并发-sync.WaitGroup
	//	WaitGroup_test()
	//	fmt.Println("WaitGroup_test end")
}
