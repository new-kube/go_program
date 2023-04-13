package main

import (
	"fmt"
	"time"

	cron "github.com/robfig/cron/v3"
)

type worker struct {
	i int
}

func NewWorker(i int) *worker {
	return &worker{
		i: i,
	}
}

func (w *worker) Work() error {
	w.i++
	fmt.Printf("worker execute time: %v, i: %v\n", time.Now(), w.i)
	return nil
}

func (w *worker) Run() {
	w.Work()
}

func main() {

	// Seconds field, required
	//c := cron.New(cron.WithSeconds())
	c := cron.New(cron.WithSeconds())

	w := NewWorker(1)

	c.AddJob("0 32 * * * *", w)
	c.Start()

	fmt.Printf("starting...\n")

	select {}

	c.Stop()
}
