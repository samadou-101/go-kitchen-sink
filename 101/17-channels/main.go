package main

import "fmt"

func main() {
	// unbuffered channel
	ch := make(chan int)

	go func() {
		ch <- 42 // send
	}()

	val := <-ch // receive
	fmt.Println(val)

	// buffered channel
	ch2 := make(chan string, 2)
	ch2 <- "a"
	ch2 <- "b"
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
