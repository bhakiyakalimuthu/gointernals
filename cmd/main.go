package main

import (
	"context"
	"fmt"
	"time"
)

func taskA(ctx context.Context, index int) {
	done := false
	go func() {
		// Keep doing something.
		for i := 0; !done; i++ {
			fmt.Printf("A%d%d\n", index, i)
		}
	}()

	// Wait util context is cancelled.
	<-ctx.Done()
	// Turn on closing flag.
	done = true
}

func main() {
	// var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	if i > 5 {
	// 		wg.Add(1)
	// 		go func(jobID int) {
	// 			defer wg.Done()
	// 			fmt.Printf("This is job: %v\n", jobID)
	// 		}(i)
	// 	} else if i < 5 {
	// 		wg.Add(1)
	// 		go func(jobID int) {
	// 			defer wg.Done()
	// 			fmt.Printf("This is job: %v\n", jobID)
	// 		}(i)
	// 	}
	// }
	// wg.Wait()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	JobWithCtx(ctx, 10)
	cancel()
}

func JobWithCtx(ctx context.Context, jobID int) {
	// tick := time.Tick(500 * time.Millisecond)
	i := 0
	for {
		if i == 0 {
			nCtx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
			keepDoingSomething(nCtx)
			cancel()
		}
		select {
		case <-ctx.Done():
			fmt.Printf("context cancelled job %v terminting\n", jobID)
			return
		default:
		}
		i++
	}
}

func keepDoingSomething(ctx context.Context) {
	// timeout := time.After(60 * time.Second)
	// tick := time.Tick(500 * time.Millisecond)
	// // Keep trying until we're timed out or got a result or got an error
	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-ctx.Done():
			fmt.Println("timeout from keepDoingSomething")
			return
		// Got a tick, we should check on doSomething()
		default:
			fmt.Println("keepDoingSomething")

			// doSomething() didn't work yet, but it didn't fail, so let's try again
			// this will exit up to the for loop
		}
		time.Sleep(5 * time.Second)
	}
}
