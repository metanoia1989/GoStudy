package main

import (
	"fmt"
	"math/cmplx"
)

/**
 * 变量定义
 */
func var_define() {
	var a int
	var b = 1
	message := "hello world"
	fmt.Println(a, b, message)

	var v1, v2 int = 5, 6
	fmt.Println(v1, v2)
}

/**
 * 数据类型
 */
func simple_datatype() {
	var a bool = true
	var b int = 1
	var c string = "hello world"
	var d float32 = 1.222
	var x complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println(a, b, c, d, x)

}

func arr_and_slice() {
	// 数组及切片定义
	var arr [5]int
	var multiD [2][3]int
	// 切片是数组的抽象。 切片使用数组作为底层结构。 切片包含三个组件：容量，长度和指向底层数组的指针
	var slice []int               // 没有容量限制的切片
	numbers := make([]int, 5, 10) // 初始长度为5，容量为10的切片
	fmt.Println(arr, multiD, slice, numbers)

	// 通过使用 append 或 copy 函数可以增加切片的容量
	// append 函数可以为数组的末尾增加值，并在需要时增加容量。
	numbers = append(numbers, 1, 2, 3, 4)
	// 增加切片容量的另一种方法是使用复制功能。 只需创建另一个具有更大容量的切片，并将原始切片复制到新创建的切片
	numbers2 := make([]int, 15)
	copy(numbers2, numbers)
	fmt.Println(numbers, numbers2)

	// 创建切片的子切片
	numbers3 := []int{1, 2, 3, 4}
	fmt.Println("numbers: ", numbers)
	slice1 := numbers3[2:]
	fmt.Println("numbers[2:]: ", slice1)
	slice2 := numbers3[:3]
	fmt.Println("numbers[:3]: ", slice2)
	slice3 := numbers3[1:4]
	fmt.Println("numbers[1:4]: ", slice3)
}

func map_type() {
	fmt.Println("map 映射：")
	// Key 类型为 string、Value 类型为 int 的 map 类型的变量
	m := make(map[string]int)
	m["clearity"] = 2
	m["simplicity"] = 3
	fmt.Println(m["simplicity"])
	fmt.Println(m["clearity"])
}

/**
 * 类型转换
 */
func type_convert() {
	fmt.Println("类型转换：")
	a := 1.1
	b := int(a)
	fmt.Println(b)
}

/**
 * 流程控制
 */
func flow_control() {
	fmt.Println("流程控制：")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("none")
	}

	fmt.Println("for循环：")
	i = 0
	sum := 0
	for i < 10 {
		sum += 1
		i++
	}
	fmt.Println(sum)

	sum = 0
	for i = 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	i = 20
	for {
		i--
		fmt.Print(i, ",")
		if i == 0 {
			break
		}
	}
	fmt.Println("")
}

func increment(i *int) {
	*i++
}

/**
 * 指针演示
 */
func pointer() {
	fmt.Println("指针演示：")
	var ap *int
	a := 12
	ap = &a
	fmt.Println(ap, *ap)

	i := 10
	increment(&i)
	fmt.Println(i)
}

/**
 * 多返回值
 */
func add(a int, b int) (int, string) {
	c := a + b
	return c, "successfullly added"
}

func main() {
	var_define()
	simple_datatype()
	arr_and_slice()
	map_type()

	type_convert()
	flow_control()
	pointer()

	sum, message := add(2, 1)
	fmt.Println(message)
	fmt.Println(sum)
}
