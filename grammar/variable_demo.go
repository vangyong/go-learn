package grammar

import "fmt"

//常量 用驼峰命名
const (
	utimePos = 13
	stimePos = 14
)

//基本类型变量
func printBaseVariable() {
	var (
		userName  string = "xiaoxiao"
		age       uint32
		high      int64
		address   string
		foreigner bool
	)
	age = 20
	high = 120
	address = "成都市 双流区"
	foreigner = true
	fmt.Println(userName)
	fmt.Println(age)
	fmt.Println(high)
	fmt.Println(address)
	fmt.Println(foreigner)
	fmt.Println("const 静态变量:")
	fmt.Println(stimePos)
}

//数组
//函数 1.使用驼峰命名 2.如果包外不需要访问用小写开头的函数 3.包外需要访问使用大写开头
func printArray() {
	var balance = [5]float32{200.0, 78.1}
	fmt.Println(balance)
	fmt.Printf("len:%d cap:%d slice=%v\n", len(balance), cap(balance), balance)
}

//指针
func printPointer() {
	var a3 int = 20
	var ip *int
	ip = &a3
	fmt.Printf("a3 的变量地址是：%x\n", &a3)
	fmt.Printf("ip的地址是：%x\n", ip)
	fmt.Printf("*ip的值：%x\n", *ip)
}

//引用类型变量和值类型变量
func printVariableDifferent() {
	var a1 string = "abcd"
	var a2 string = a1
	fmt.Println(&a1)
	fmt.Println(&a2)

	var a3 int = 12001
	var a4 int = a3
	fmt.Println(&a3)
	fmt.Println(&a4)
}
