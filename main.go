package main

import (
	"fmt"
	"time"
)

func main1(){
	var c chan int
	bc := make(chan int)
	fmt.Println(c)
	fmt.Println(bc)
}

func greet(c chan string)  {
	fmt.Println("greeting from  : ",  <- c)
}
func main2()  {
	fmt.Println("main started")
	c:= make(chan string)

	go greet(c)

	c <- "greet"
	close(c)
	c <- "greeting"
	fmt.Println("main ended")

}

func square( c chan int){
	for i := 1; i < 9; i++{
		c <- i*i
		fmt.Println("square running ", i)
	}
	close(c)
}
func main3()  {
	fmt.Println("main started")
	c:= make(chan int)
	go square(c)

	for {
		val, ok := <- c
		if ok == false {
			fmt.Println(val, ok)
			break
		}else {
			fmt.Println(val, ok)
		}
	}
	fmt.Println("main stopped")
}
func main4()  {
	fmt.Println("main started")
	c:= make(chan int)
	go square(c)

	for val := range c {
		fmt.Println(val)
	}
	fmt.Println("main stopped")
}

func squares( c chan int){
	for i := 0; i <3; i++{
		num := <- c
		fmt.Println("square running ", num)
	}
	close(c)
}
func main5()  {
	fmt.Println("main started")
	c:= make(chan int,3)
	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	fmt.Println("main stopped")
}

type order struct {
	name string
}
func food (c chan []order) {
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

func kitchen () {
	fmt.Println("running kitchen")
	time.Sleep(10)
}

func drone () {
	fmt.Println("running drone")
	time.Sleep(10)
}

func main() {
	fmt.Println("main started")
	c:= make(chan []order,1)
	go food(c)
	orders := <- c
	// close(c)
	for _,order := range orders {
		if order.name == "new" {
			fmt.Println(order.name)
			go kitchen()
		}else if order.name == "preparation" {
			fmt.Println(order.name)
			go drone()
		}
	}
	fmt.Println("main stopped")
}