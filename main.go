package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	var c chan int
	bc := make(chan int)
	fmt.Println(c)
	fmt.Println(bc)
}

func greet(c chan string) {
	fmt.Println("greeting from  : ", <-c)
}
func main2() {
	fmt.Println("main started")
	c := make(chan string)

	go greet(c)

	c <- "greet"
	close(c)
	c <- "greeting"
	fmt.Println("main ended")

}

func square(c chan int) {
	for i := 1; i < 9; i++ {
		c <- i * i
		fmt.Println("square running ", i)
	}
	close(c)
}
func main3() {
	fmt.Println("main started")
	c := make(chan int)
	go square(c)

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok)
			break
		} else {
			fmt.Println(val, ok)
		}
	}
	fmt.Println("main stopped")
}
func main4() {
	fmt.Println("main started")
	c := make(chan int)
	go square(c)

	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main stopped")
}

func squares(c chan int) {
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("square running ", num)
	}
	close(c)
}
func main5() {
	fmt.Println("main started")
	c := make(chan int, 3)
	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	fmt.Println("main stopped")
}

type order struct {
	name string
}

func food(c chan []order) {
	time.Sleep(10)
	c <- []order{
		{
			"new",
		},
		{
			"preparation",
		},
	}
	close(c)
}

func kitchen() {
	fmt.Println("running kitchen")
	time.Sleep(10)
}

func drone() {
	fmt.Println("running drone")
	time.Sleep(10)
}

func main6() {
	fmt.Println("main started")
	c := make(chan []order, 1)
	go food(c)
	orders := <-c
	// close(c)
	for _, order := range orders {
		if order.name == "new" {
			fmt.Println(order.name)
			go kitchen()
		} else if order.name == "preparation" {
			fmt.Println(order.name)
			go drone()
		}
	}
	fmt.Println("main stopped")
}

type Point struct {
	Lon string
	Lat string
}

func recv(c1 chan Point) {
	fmt.Println("recv started")
	r1 := <-c1
	fmt.Println("r1", r1)
	r2 := <-c1
	fmt.Println("r2", r2)
	r3 := <-c1
	fmt.Println("r3", r3)
	// if r2.Lon == "2" {
	// 	fmt.Println("inside recv if", r2.Lon)
	// 	c2 <- true
	// }
}
func recv1(c1 chan Point, c2 chan string) {
	// for {
	// 	select {
	// 	case _ = <-c3:
	// 		close(c2)
	// 		return
	// 	default:
	// 		fmt.Println("running default case")
	//
	// 	}
	// }
	for r1 := range c1 {
		fmt.Println("recv started")
		fmt.Println("r1", r1)
		if r1.Lon == "1" {
			fmt.Println("recv ended")
			c2 <- r1.Lon
			// close(c2)
			time.Sleep(2 * time.Second)
		}
	}
	close(c2)
}

func main() {
	fmt.Println("main started")
	p1 := []Point{{Lon: "1", Lat: "1"}, {Lon: "2", Lat: "2"}, {Lon: "3", Lat: "3"}, {Lon: "4", Lat: "4"}, {Lon: "5", Lat: "5"}, {}, {Lon: "3", Lat: "3"}}
	p2 := []Point{{Lon: "1", Lat: "1"}, {Lon: "2", Lat: "2"}, {Lon: "3", Lat: "3"}, {Lon: "4", Lat: "4"}, {Lon: "5", Lat: "5"}, {}, {Lon: "3", Lat: "3"}}
	p3 := [][]Point{p1, p2}
	var wg sync.WaitGroup
	wg.Add(2)
	for _, item := range p3 {
		go func() {
			defer wg.Done()
			drones(item)
		}()
	}
	wg.Wait()
}
func drones(p []Point) {

	r := make(chan string)
	c := make(chan Point, 3)
	// var wg sync.WaitGroup
	// wg.Add(1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range r {
			fmt.Println("final result true", result)
		}
	}()

	go func() {
		// defer wg.Done()
		recv1(c, r)
	}()
	counter := 0
	for _, i := range p {
		fmt.Println("inside main if", i)
		c <- i
		if i.Lat == "4" {
			fmt.Println("condition true", i.Lat)
			counter++
			// wg.Done()
			break
		}
	}
	close(c)

	if counter == 0 {
		return
	}
	wg.Wait()
}
