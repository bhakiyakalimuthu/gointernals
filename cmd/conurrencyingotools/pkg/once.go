package pkg

import (
	"fmt"
	"sync"
)

func OnceTest() {

	var wg sync.WaitGroup
	var once sync.Once
	var count int
	increment := func() {
		count++
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
