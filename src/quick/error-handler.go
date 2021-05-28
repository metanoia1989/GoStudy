package main

import (
	"errors"
	"fmt"
	"net/http"
)

func increment(n int) (int, error) {
	if n < 0 {
		// 返回一个 error 对象
		return -1, errors.New("math: cannot process negative number")
	}
	return (n + 1), nil
}

func main() {
	resp, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

	// 获取函数自定义错误问题
	num := 5
	if inc, err := increment(num); err != nil {
		fmt.Printf("Failed Number: %v, error message: %v", num, err)
	} else {
		fmt.Printf("Incrementd Number: %v", inc)
	}
}
