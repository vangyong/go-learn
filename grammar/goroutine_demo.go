package grammar

import (
	"fmt"
	"time"
)

func GoroutineSay(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine:"+ s)
	}
}
