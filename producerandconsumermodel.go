package main

import (
	"fmt"
	"time"
)

//chan <- 只能写，不能读
func producer(value chan<- int) {

	defer close(value)

	for i := 0; i < 10; i++ {
		value <- i
	}

}

//chan <- 只能读，不能写
func consumer(value <-chan int) {
	for data := range value {
		fmt.Println(data)
	}
}
func main() {

	c := make(chan int, 5)
	go producer(c)
	go consumer(c)
	time.Sleep(time.Second * 2)
}
