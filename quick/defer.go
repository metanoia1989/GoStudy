package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Retruned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Paniking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func main() {
	f()
	fmt.Println("Returned normally from f.")
}
