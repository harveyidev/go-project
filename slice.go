package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

//slice本身没有数据，是对底层array的一个view
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 =", s1)
	s2 := arr[:]
	fmt.Println("s2 =", s2)

	fmt.Println("after updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("after updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("ReSlice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)
}
