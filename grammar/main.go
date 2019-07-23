package main

import "fmt"

//常量 用驼峰命名
const (
	utimePos = 13
	stimePos = 14
	startPos = 21
	vssPos   = 22
	rssPos   = 23
)

var (
	userName    string
	age         int32
	lastTotal   int64
	lastSeconds int64
	ipcpu       int64
)

func main() {
	//slice切片
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	//循环遍历计算
	rangeSum()
}

//函数 1.使用驼峰命名 2.如果包外不需要访问用小写开头的函数 3.包外需要访问使用大写开头
func printSlice(x []int) {
	fmt.Printf("len:%d cap:%d slice=%v\n", len(x), cap(x), x)
}

func rangeSum() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}

//接口命名
// type helloWorld interface {
//     func Hello();
// }

// type SayHello helloWorld

//receiver 命名
// type A struct{}

// func (a *A) methodA() {
// }
// func (a *A) methodB() {
//     a.methodA()
// }
