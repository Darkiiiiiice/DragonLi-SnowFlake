package main

import (
	"fmt"
	"io/mariomang/github/dao"
	"io/mariomang/github/snowflake"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("BenchMark")

	var waitGroup sync.WaitGroup

	snow := snowflake.NewSnowFlake(1, 1)
	ch := make(chan int64, 1024)

	waitGroup.Add(1)
	go savehandle(&waitGroup, ch)

	s := time.Now().UnixNano() / 1e6
	for {

		e := time.Now().UnixNano() / 1e6

		if e-s >= 1 {
			// ch <- -3
			// fmt.Println("=================================")
			break
		}

		waitGroup.Add(1)
		go idhandle(&waitGroup, snow, ch)
	}

	waitGroup.Wait()
}

func idhandle(wg *sync.WaitGroup, snow *snowflake.SnowFlake, ch chan int64) {
	defer wg.Done()

	id := snow.GetID()
	// fmt.Printf("Tag-%d    ID:%d\n", tag, id)
	ch <- id

}

func savehandle(wg *sync.WaitGroup, ch chan int64) {
	defer wg.Done()

	for {
		select {
		case id := <-ch:
			fmt.Println(wg)
			fmt.Printf("goroutine-%d   Saved-%d\n", runtime.NumGoroutine(), id)
			dao.InsertByWorkID(id, "TEST")
		}

		// if id == -3 {
		// 	fmt.Println("Os.Exited()")
		// 	break
		// }

		// select {
		// case id := <-ch:
		// 	fmt.Printf("IDD:%d\n", id)
		// }
	}
}
