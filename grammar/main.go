package main

import "fmt"

func main() {
	//slice切片
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	//循环遍历计算
	rangeSum()
}

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
