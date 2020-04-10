package main

import (
	"fmt"
	"time"
)

func worker(id int, done chan bool) {
	fmt.Printf("%d working...\n", id)
	time.Sleep(time.Second)
	fmt.Printf("%d done\n", id)

	done <- true
}

func main() {
	count := 5

	// sync
	done := make(chan bool, 1)

	// make 5 concurrent go routine
	for i := 0; i < count; i++ {
		go worker(i, done)
	}

	// pop from channel the require count goroutine needs to push to channel
	for i := 0; i < count; i++ {
		<-done
	}
}

// We can use channels to synchronize execution across goroutines.
// Here’s an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
