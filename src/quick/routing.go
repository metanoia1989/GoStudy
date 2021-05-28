package main

import (
	"fmt"
	"time"
)

func c() {
	time.Sleep(time.Second * 2)
	fmt.Println("I am concurrent")
}

func main() {
	go c()
	fmt.Println("I am main")
	time.Sleep(time.Second * 3)
}
