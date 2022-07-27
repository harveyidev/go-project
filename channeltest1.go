package main

import (
	"fmt"
	"time"
)

func main() {

	//channel的定义格式
	//c := make(chan, 数据类型)无缓冲的channel
	c := make(chan int, 0)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 5; i++ {
		value := <-c
		fmt.Println(value)
	}

	time.Sleep(time.Second)
}
