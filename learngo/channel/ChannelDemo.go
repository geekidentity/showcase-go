package main

import (
	"fmt"
	"time"
)

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
		close(channels[i])
	}
	time.Sleep(time.Second)
}

func createWorker(id int) chan<- int {
	channel := make(chan int)
	go func() {
		for n := range channel {
			fmt.Printf("worker %d received %c\n", id, n)
		}
	}()
	return channel
}

func worker(i int, c chan int) {
	for {
		n, ok := <-c
		if ok {
			fmt.Printf("worker %d received %c\n", i, n)
		} else {
			break
		}
	}
}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
}

func channelClose() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

func main() {
	channelDemo()
	bufferedChannel()
}
