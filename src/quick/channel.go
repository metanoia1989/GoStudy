package main

import "fmt"

// 用 channel 与协程进行通信
func main() {
	c := make(chan string)
	go func() {
		c <- "hello"
	}()
	msg := <-c
	fmt.Println(msg)
}
