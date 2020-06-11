package pkg

import (
	"fmt"
	"time"
)

func PreventMemLeak() {
	dowork := func(strings chan string, done <-chan interface{}) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Printf("do work exited \n")
			defer close(terminated)
			for {
				select {
				case <-done:
					fmt.Println("do work done \n ")
					return
				case s := <-strings:
					fmt.Println(s)
					close(strings)
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := dowork(nil, done)
	go func() {
		time.Sleep(time.Second * 5)
		close(done)
	}()

	<-terminated
	fmt.Printf("Done \n")
}
