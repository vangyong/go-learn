package grammar

func ChannelSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum  //把 sum 发送到通道 c
}
