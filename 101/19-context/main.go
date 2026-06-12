package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// background context (root)
	ctx := context.Background()
	fmt.Println(ctx.Err())

	// cancelable context
	ctx2, cancel := context.WithCancel(ctx)
	cancel()
	fmt.Println(ctx2.Err())

	// timeout context
	ctx3, cancel2 := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel2()
	<-ctx3.Done()
	fmt.Println(ctx3.Err())
}
