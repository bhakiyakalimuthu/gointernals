package main

import (
	"fmt"
	"sync"

	"github.com/bhakiyakalimuthu/gointernal/cmd/conurrencyingotools/pkg"
)

func main() {
	// one()
	// two()
	// three()

	// pkg.CondTest()
	// pkg.OnceTest()
	pkg.PreventMemLeak()
}

func one() {
	salutation := "hello"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	fmt.Println("\n \n")
	fmt.Println(salutation)
	wg.Wait()
}

func two() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func three() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
