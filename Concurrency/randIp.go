package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	now := time.Now()
	ipTingkat1 := make(chan string)
	ipTingkat2 := make(chan string)
	ipTingkat3 := make(chan string)
	ipTingkat4 := make(chan string)
	buf := make([]byte, 4)

	var ws sync.WaitGroup

	ws.Add(4)

	go func() {
		defer ws.Done()
		go generateIp(1, buf, ipTingkat1)
		for {
			ip, open := <-ipTingkat1

			if !open {
				break
			}
			fmt.Println("From Thread 1 :", ip)
		}
	}()

	go func() {
		defer ws.Done()
		go generateIp(2, buf, ipTingkat2)
		for {
			ip, open := <-ipTingkat2

			if !open {
				break
			}
			fmt.Println("From Thread 2 :", ip)
		}
	}()

	go func() {
		defer ws.Done()
		go generateIp(3, buf, ipTingkat3)
		for {
			ip, open := <-ipTingkat3

			if !open {
				break
			}
			fmt.Println("From Thread 3 :", ip)
		}
	}()

	go func() {
		defer ws.Done()
		go generateIp(4, buf, ipTingkat4)
		for {
			ip, open := <-ipTingkat4

			if !open {
				break
			}
			fmt.Println("From Thread 4 :", ip)
		}
	}()

	ws.Wait()
	fmt.Println("Time taken :", time.Since(now))

}

func generateIp(tingkat int64, buf []byte, out chan string) {
	defer close(out)
	for i := 0; i < 20; i++ {
		rand.Seed(tingkat * time.Now().UnixNano())
		ip := rand.Uint32()
		binary.LittleEndian.PutUint32(buf, ip)
		out <- net.IP(buf).String()
		time.Sleep(time.Second / 2)
	}
}
