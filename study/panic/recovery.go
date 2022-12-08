package main

import (
	"fmt"
	"sync"
)

func main() {
	recovery := func() {
		err := recover()
		if err != nil {
			fmt.Printf("go routine panic: %+v\n", err)
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer recovery()
		defer wg.Done()

		fmt.Printf("hello, i'm ok...\n")
		panic("hello, i'm craze...\n")
	}(&wg)

	wg.Wait()
	fmt.Printf("safety exit...\n")
}
