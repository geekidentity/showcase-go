package main

import (
	"fmt"
	"sync"
)

func channelDemoWg() {
	var workers [10]workerWg
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorkerWg(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}
	wg.Wait()
	// wait for all of them

}

type workerWg struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorkerWg(id int, wg *sync.WaitGroup) workerWg {
	w := workerWg{
		in: make(chan int),
		wg: wg,
	}
	go doWorkerWg(id, w.in, w.wg)
	return w
}

func doWorkerWg(id int, channel chan int, wg *sync.WaitGroup) {
	for n := range channel {
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done()
	}
}
