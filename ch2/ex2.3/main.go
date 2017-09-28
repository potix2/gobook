package main

import (
	"fmt"
	"github.com/potix2/gobook/ch2/popcount"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		popcount.PopCount(1111)
	}
	elapsed := time.Since(start)
	fmt.Printf("Popcount took %s\n", elapsed)

	start = time.Now()
	for i := 0; i < 100000; i++ {
		popcount.PopCountLoop(1111)
	}
	fmt.Printf("PopCountLoop took %s\n", time.Since(start))
}
