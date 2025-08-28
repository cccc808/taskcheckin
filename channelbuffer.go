// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	fmt.Println("producer strat")
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("send :", i)
	}

	fmt.Println("producer done")
}

func consumer(ch chan int) {
	fmt.Println("consumer start:")
	for o := range ch {
		fmt.Println("receive :", o)
	}

	fmt.Println("consumer done")
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Second)
	fmt.Println("main done")
}
