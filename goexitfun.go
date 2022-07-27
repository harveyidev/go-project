package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//匿名函数
	go func() {
		//延迟调用
		defer fmt.Println("hello world")
		func() {
			fmt.Println("hello world 111")
			//退出当前Goroutine
			runtime.Goexit()
			fmt.Println("hello world 222")
		}()
	}()

	time.Sleep(10 * time.Second)

}
