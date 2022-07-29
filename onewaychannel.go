package main

import (
	"fmt"
	"time"
)

func main01() {

	c := make(chan int, 5)

	//赋值给单向的chan
	var send chan<- int = c
	var recv <-chan int = c

	send <- 100 //ok
	//value := <-send //err
	value := <-recv //ok
	//recv := <-100 //err

	fmt.Println(value)

	//不能将单向的chan转成双向的
	//v1 :=chan int(send)//err
	//v2 := chan int(recv)//err
}

//chan<-  写
func counter(value chan<- int) {
	for i := 0; i < 10; i++ {
		value <- i
	}
	close(value)
}

//<-chan  读
func printer(value <-chan int) {
	for data := range value {
		fmt.Println(data)
	}
}

func main() {
	c := make(chan int)
	go counter(c) //生产者
	go printer(c) //消费者
	time.Sleep(time.Second)
}
