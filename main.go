package main

import (
	"fmt"
	"go-learn/basic"
)

func main() {

	// 基本数据类型
	//basic.PrintBaseVariable()

	//引用类型和值类型
	//basic.PrintDifferentVariable()
	//fmt.Println("PrintDifferentVariable end")

	//数组
	//basic.PrintArray()

	//结构
	//var book1 basic.Book
	//book1.Author = ""
	//book1.Author = "www.runoob.com"
	//book1.Subject = "Go 语言教程"
	//book1.BookId = 6495407
	//basic.PrintStruct(book1)

	//循环遍历计算
	//basic.SumRange()

	//切片(自动增长缩小)
	//var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	//var sum = basic.SumSlice(numbers)
	//fmt.Println(sum)

	//接口
	var connector basic.Connector
	//connector = new (basic.PhoneConnector)
	connector = basic.PhoneConnector{"telephone"}
	connectorName := connector.SetName("IPhone4 ")
	fmt.Println("connectorName:"+connectorName)
	fmt.Println("-----------")
	fmt.Println(connector.GetName())
	fmt.Println("-----------:")
	connector.Connect()

	//c:=basic.Cube{3}
	//var s basic.Shape = c
	//var o basic.Object= c
	//fmt.Println("volume of s of interface type shape is",s.Area())
	//fmt.Println("area of o of interface type object is",o.Volume())

	//map
	//basic.RangePrintMap()

	//并发-goroutine轻量级线程
	//go basic.GoroutineSay("world")
	//basic.GoroutineSay("hello")

	//并发-channel通道,类似于栈，先进后出
	//ch <- v  把v发送到通道ch，v := <-ch  从ch接收数据并把值赋给v
	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go basic.ChannelSum(s[:len(s)/2], c) //计算前一半数据，并写入通道c
	//go basic.ChannelSum(s[len(s)/2:], c) //计算后一半数据，并写入通道c
	//x, y := <-c, <-c   //从通道 c 中接收
	//fmt.Println(x, y, x+y)

	//通道缓冲区
	//ch := make(chan int, 2)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//遍历通道与关闭通道
	//c := make(chan int, 10)
	//go basic.Fibonacci(cap(c), c)
	//// range 遍历从通道接收到的数据，因为 c 在发送完 10 个数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据之后就结束了。
	//// 如果上面的 c 通道不关闭，那么 range 函数就不会结束，从而在接收第 11 个数据的时候就阻塞了。
	//for i := range c {
	//	fmt.Println(i)
	//}

	//并发-sync.WaitGroup
	// WaitGroup变量定义后，是不允许被拷贝的，即不允许作为函数参数或者赋值给其他变量
	//basic.WaitGroupTest()
	//fmt.Println("WaitGroupTest end")

	// UDP服务
	//address := "0.0.0.0:8000"
	//basic.UdpServer(address)


	// json struct map 转换
	//jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,
	//			"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/",
	//			"serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
	//map1 := basic.JsonToMap(jsonStr)
	//config := basic.JsonToStruct(jsonStr)
	//jsonStr1 := basic.StructToJson(config)
	//basic.MapToJson(map1)
	//basic.JsonToStruct(jsonStr1)

	// 数组 json 转换
	//arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	//jsonStr2 := basic.ArrayToJson(arr)
	//fmt.Println(jsonStr2)
	//arr2 := basic.JsonToArray(jsonStr2)
	//fmt.Println(arr2)

}
