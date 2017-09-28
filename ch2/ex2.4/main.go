package main

import (
	"fmt"
	"github.com/potix2/gobook/ch2/popcount"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		popcount.PopCountShift(1111)
	}
	fmt.Printf("Popcount3 took %s\n", time.Since(start))
}
