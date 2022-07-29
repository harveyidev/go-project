package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 0)

	//开启协程
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//关闭chan
		close(c)
	}()

	//1、死循环
	//for {
	//	//判断channel是否关闭
	//	//没有关闭才能读取数据
	//	if value, ok := <-c; ok {
	//		fmt.Println(value)
	//	} else {
	//		fmt.Println("chan已关闭")
	//		break
	//	}
	//}

	//2、通过range做范围遍历
	for data := range c {
		fmt.Println(data)
	}

	//sleep一秒钟
	time.Sleep(time.Second)

}
