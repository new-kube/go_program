package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // 遇见io.EOF，注意通常是Ctrl+D或者Ctrl+Z
		counts[input.Text()]++
		// 等价于:
		// line := input.Text()
		// counts[line] = counts[line] + 1
		// 当map中不存在line时，counts[line] = 0，是根据int值类型推演为0.
	}

	// 注意，忽略input.Err()中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t %s\n", n, line)
		}
		// line即key的顺序不是固定的。每次都不相同。是随机的。这是有意设计的。即不对排序做任何保证。
	}
}
