package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

var cancelBefore = false

func main() {
	// contextCancel()
	contextBackground()
}

func contextBackground() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	cancel()
	// go func() {
	// r := bufio.NewScanner(os.Stdin)
	// r.Scan()
	// 	time.Sleep(time.Second)
	// 	cancel()
	// }()
	// time.AfterFunc(time.Second, cancel)
	sleepAndTalk(ctx, time.Second*3, "hello")
}

func sleepAndTalk(ctx context.Context, duration time.Duration, str string) {
	select {
	case <-ctx.Done():
		log.Print(ctx.Err())
	case <-time.After(duration):
		fmt.Println(str)
	}

}

func contextCancel() {
	c, cCancel := context.WithTimeout(context.Background(), time.Second*3)

	defer func() {
		fmt.Println("c.ctx:", c.Err())
		cCancel()
	}()

	c1, cf1 := context.WithTimeout(c, time.Second*10)
	func(ctx context.Context) {
		fmt.Println()
		time.Sleep(time.Second * 6)
	}(c1)
	defer func() {
		fmt.Println("c1.ctx:", c.Err())
		cf1()
	}()

}
func contextChain() {
	c, cCancel := context.WithCancel(context.Background())

	c1, cf1 := context.WithCancel(c)
	defer cf1()

	c2, cf2 := context.WithCancel(c)
	defer cf2()

	c11, cf11 := context.WithCancel(c1)
	defer cf11()

	c12, cf12 := context.WithCancel(c1)
	defer cf12()

	if cancelBefore {
		cCancel()
	}

	for k, c := range map[string]context.Context{`c1`: c1, `c11`: c11, `c12`: c12, `c2`: c2} {
		var s string
		if c.Err() != nil {
			s = `cancelled`
		} else {
			s = `not cancelled`
		}
		println(k + ` is ` + s)
	}

	if !cancelBefore {
		cCancel()
	}
}
