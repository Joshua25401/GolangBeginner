package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	output := make(chan int)
	output2 := make(chan int)

	now := time.Now()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		go counting(10, output)
		for {
			item, open := <-output
			if !open {
				break
			}
			fmt.Println("From Thread 1 :", item)
		}
	}()

	go func() {
		defer wg.Done()
		go counting(20, output2)
		for {
			item, open := <-output2
			if !open {
				break
			}
			fmt.Println("From Thread 2 :", item)
		}
	}()

	wg.Wait()
	fmt.Println("Time taken :", time.Since(now))
}

func counting(item int, out chan int) {
	defer close(out)
	for i := 0; i < item; i++ {
		out <- i
		// time.Sleep(time.Second / 2)
	}
}
