package main

import "fmt"

func main() {
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

}

func printSlice(x []int) {
	fmt.Printf("len%d cap=%d slice=%v\n", len(x), cap(x), x)
}
