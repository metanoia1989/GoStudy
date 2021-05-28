package main

import "fmt"

func add(a int, b int) int {
	c := a + b
	return c
}

/**
 * 返回值预声明
 */
func bbc(a int, b int) (c int) {
	c = a + b
	return
}

func main() {
	fmt.Println(add(2, 1))
	fmt.Println(bbc(5, 6))
}
