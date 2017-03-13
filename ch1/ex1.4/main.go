package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		counts = countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
				continue
			}
			tmp := countLines(f)
			f.Close()

			// merge counts
			for line, n := range tmp {
				counts[line] += n
				if _, ok := filenames[line]; !ok {
					filenames[line] = make(map[string]bool)
				}
				filenames[line][arg] = true
			}
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t", n)
			for k, _ := range filenames[line] {
				fmt.Printf("%s,", k)
			}
			fmt.Printf("\n%s\n", line)
		}
	}
}

func countLines(f *os.File) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	return counts
}
