package grammar

import (
	"fmt"
	"runtime"
	"sync"
)

func SyncWaitGroupTest() {
	fmt.Println("cpus:" )
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go GoSyncTest(&wg, i)
	}
	wg.Wait()
}

func GoSyncTest(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 1000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done()
}
