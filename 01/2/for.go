package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		go func() {
			fmt.Println(arg)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("---------------")

	for _, arg := range os.Args[1:] {
		go func(a string) {
			fmt.Println(a)
		}(arg)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("---------------")

	for _, arg := range os.Args[1:] {
		a := arg
		go func() {
			fmt.Println(a)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("---------------")
}
