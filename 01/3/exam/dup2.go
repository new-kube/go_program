// 打印重复的行，并打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileLines := make(map[string]map[string]int)

	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2 failed: %v\n", err)
			continue
		}
		fileLines[file] = make(map[string]int)
		input := bufio.NewScanner(f)
		for input.Scan() {
			line := input.Text()
			fileLines[file][line]++
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			files := make([]string, 1)
			for file, lines := range fileLines {
				if lines[line] > 0 {
					files = append(files, file)
				}
			}
			fmt.Printf("%d\t%s\t%v\n", n, line, files)
		}
	}
}
