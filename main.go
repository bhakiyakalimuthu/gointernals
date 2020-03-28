package main

import "fmt"

func main(){
	var c chan int
	bc := make(chan int)
	fmt.Println(c)
	fmt.Println(bc)
}

