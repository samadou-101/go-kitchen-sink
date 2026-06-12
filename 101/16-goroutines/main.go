package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// goroutine = lightweight thread
	go say("hello")
	go say("world")

	// give goroutines time to run
	time.Sleep(100 * time.Millisecond)
}
