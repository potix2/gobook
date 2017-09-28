package main

import (
	"fmt"
	"time"
)

func PopCount(x uint64) int {
	count := 0
	return count
}

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		PopCount(1111)
	}
	fmt.Printf("Popcount took %s\n", time.Since(start))
}
