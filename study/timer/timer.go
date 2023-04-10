package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func doWork(ctx context.Context) (time.Duration, error) {

	begin := time.Now()

	x := rand.Intn(10)
	fmt.Printf("x = %d\n", x)
	time.Sleep(time.Second * time.Duration(x))

	elapsed := time.Now().Sub(begin)

	fmt.Printf("scheduler, doWork usage(%dms)\n", elapsed/time.Second)
	return elapsed, nil
}

func doSched(ctx context.Context) error {
	interval := time.Second * 10
	left := interval

	timer := time.NewTimer(interval)
	for {
		fmt.Printf("\n-----next loop, now=%s\n", time.Now().Format("2006-01-02 15:04:05"))
		select {
		case <-timer.C:
			fmt.Printf("begin work...now=%s\n", time.Now().Format("2006-01-02 15:04:05"))
			elapsed, _ := doWork(ctx)
			left = interval - elapsed
			if left < 0 {
				left = 0
			}
			fmt.Printf("end work: usage=%ds, left=%d...\n", elapsed/time.Second, left/time.Second)
			break
		case <-ctx.Done():
			fmt.Printf("doSched exit...\n")
			return nil

		}
		fmt.Printf("reset=%d\n", left/time.Second)
		timer.Reset(left)
	}
}

func main() {
	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)

	go doSched(ctx)

	time.Sleep(time.Second * 200)

	ctx.Done()
	time.Sleep(time.Second)

	fmt.Printf("exit...\n")
}
