package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

// 方法是一个特殊类型的带有返回值的函数。返回值既可以是值，也可以是指针
func (p *person) describe() {
	fmt.Printf("%v is %v years old.\n", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p person) setName(name string) {
	p.name = name
}

func main() {
	// 实例化结构体
	p := person{name: "Bob", age: 42, gender: "Male"}
	p2 := person{"Bob", 42, "Male"}

	fmt.Println(p.name, p.age, p.gender)

	var pp *person
	pp = &p2
	fmt.Println(pp.name, pp.age, pp.gender)

	// 调用方法
	pp.describe()
	pp.setAge(45)
	fmt.Println(pp.name, pp.age)
	pp.setName("Hari") // 无效的更改
	fmt.Println(pp.name, pp.age)
}
