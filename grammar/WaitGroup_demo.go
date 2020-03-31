package grammar

import (
	"fmt"
	"runtime"
	"sync"
)

func WaitGroupTest() {
	fmt.Println(runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go GoTest(&wg, i)
	}
	wg.Wait()
}

func GoTest(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
