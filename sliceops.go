package main

import "fmt"

func main() {
	var s []int
	s = nil
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}
