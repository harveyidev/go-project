package main

import (
	"fmt"
	"time"
)

func main() {

	//有缓冲的
	c := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("长度", len(c), "容量", cap(c))
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		value := <-c
		fmt.Println(value)
	}
}
