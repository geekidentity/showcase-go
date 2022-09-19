package main

import (
	"fmt"
)

func channelDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func doWorker(id int, channel chan int, done chan bool) {
	for n := range channel {
		fmt.Printf("worker %d received %c\n", id, n)
		go func() {
			done <- true
		}()
	}
}

func main() {
	//channelDemo()
	//channelDemoWg()
}
