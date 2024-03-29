package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""

	begin := time.Now()

	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}

	elapsed := time.Since(begin)

	begin = time.Now()

	s = strings.Join(os.Args, " ")

	elapsed2 := time.Since(begin)

	fmt.Printf("+ time: %d, join time: %d\n", elapsed, elapsed2)
	// time.Duration的单位是 Nanosecond
	// 1s = 1000 millisecond
	// 1ms = 1000 microsecond
	// 1us = 1000 nanosecond
	// 1ns = 1 nanosecond
}
