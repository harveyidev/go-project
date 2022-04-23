package main

import "fmt"

//指针不能运算
func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}
