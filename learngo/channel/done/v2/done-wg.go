package v2

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
	in   chan int
	done func()
}

func createWorkerWg(id int, wg *sync.WaitGroup) workerWg {
	w := workerWg{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorkerWg(id, w)
	return w
}

func doWorkerWg(id int, w workerWg) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}
}
