// Dup1は標準入力から2回以上現れる行を出現回数と一緒に表示します。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// 注意: input.Err()からの可能性を無視している
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
