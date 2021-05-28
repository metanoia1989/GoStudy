package main

import (
	"custom_package/person"
	"fmt"
)

func main() {
	p := person.Description("Mailp")
	fmt.Println(p)
}
