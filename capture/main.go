package main

import "fmt"
import "io/ioutil"
import "goshark"
import "golibwireshark"

func main() {
	// var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// printSlice(numbers)

	// nums := []int{2, 3, 4}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("sum:", sum)

	file := "E:\testfile\2019-04-23 14_12_45-HTTP.pcap"
	d := goshark.NewDecoder()
	if err := d.DecodeStart(file); err != nil {
		log.Println("Decode start fail:", err)
		return
	}
	defer d.DecodeEnd()

	f, err := d.NextPacket()
	if err != nil {
		log.Println("Get packet fail:", err)
		return
	}

	key := "igmp.maddr"
	value, ok := f.Iskey(key)
	if ok {
		fmt.Printf("key: %s\nvalue: %s\n", key, value)
	}
}

//打印数组
func printSlice(x []int) {
	fmt.Printf("len%d cap=%d slice=%v\n", len(x), cap(x), x)
}
