package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"time"
	// "time"
)

func main() {

	buf := make([]byte, 4)

	for i := 0; i < 10; i++ {

		rand.Seed(1 * time.Now().UnixNano())
		ip := rand.Uint32()
		binary.LittleEndian.PutUint32(buf, ip)
		fmt.Printf("%s\n", net.IP(buf))
	}
}
