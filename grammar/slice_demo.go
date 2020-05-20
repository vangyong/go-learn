package grammar

import (
	"fmt"
)

func SumSlice(nums []int) int {
	fmt.Printf("len%d cap=%d slice=%v\n", len(nums), cap(nums), nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
