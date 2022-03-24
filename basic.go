package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//
//var aa = 3
//var ss = "kkk"
//var bb = true

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variable() {
	var a int = 0
	var s string
	fmt.Printf("%d %q \n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//只能在函数内部使用
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func euler1() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

func consts() {
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	//b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	//fmt.Println("hello world")
	//variable()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	fmt.Println(aa, bb, ss)
	euler()
	euler1()
	consts()
	enums()
}
