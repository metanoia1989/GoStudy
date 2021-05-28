package main

import "fmt"

func sc(ch chan<- string) {
	ch <- "hello"
}

// 单向通道
func main() {
	ch := make(chan string)

	go sc(ch)
	fmt.Println(<-ch)

}
