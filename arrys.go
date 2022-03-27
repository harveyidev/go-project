package main

import "fmt"

//值传递，拷贝数组
func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v) //索引和值
	}
}

//指针
func printPArray(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v) //索引和值
	}
}
func main() {

	var arr1 [5]int

	arr2 := [3]int{1, 3, 5}

	arr3 := [...]int{2, 4, 6, 8, 10}

	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i := range arr3 {
		fmt.Println(arr3[i]) //同上
	}

	for i, v := range arr3 {
		fmt.Println(i, v) //索引和值
	}
	printArray(arr1)
	//printArray(arr2)
	printArray(arr3)
}
