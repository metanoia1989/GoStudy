package main

import (
	"fmt"
	"time"
)

// select 多通道选择

func speed1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "speed 1"
}

func speed2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "speed 2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go speed1(c1)
	go speed2(c2)

	select {
	case s1 := <-c1:
		fmt.Println(s1)
	case s2 := <-c2:
		fmt.Println(s2)
	}
}
