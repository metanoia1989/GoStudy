package main

import "fmt"

// 缓冲通道，只允许接收指定个数的值
func main() {
	ch := make(chan string, 5)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch2 := make(chan int, 1)
	ch2 <- 1
	ch2 <- 2 // 报错
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
