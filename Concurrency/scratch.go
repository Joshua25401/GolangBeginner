package main

import (
	"fmt"
	"sync"
	"time"
)

func task() {
	for i := 0; i < 10; i++ {
		fmt.Println("Task", i)
		time.Sleep(time.Second)
	}
}

func main() {
	now := time.Now()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Scanln()
		go task()
		go task()
		go task()
	}()

	wg.Wait()
	fmt.Println("Time taken :", time.Since(now))

}
