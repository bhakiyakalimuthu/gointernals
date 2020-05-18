package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	t1, _ := time.Parse(layout, "2011-03-22 08:10:00")
	t2, _ := time.Parse(layout, "2011-03-22 08:11:29")

	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.After(t1))
	// fmt.Println(t.Unix())
}
