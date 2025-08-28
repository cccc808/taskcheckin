// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全
package main

import (
	"fmt"
	"sync"
	"time"
)

func addnum(ch chan int, counter *int, mu *sync.Mutex) {
	mu.Lock()
	for i := 0; i < 1000; i++ {
		(*counter)++
	}

	fmt.Println("counter is:", *counter)

	mu.Unlock()
}

func main() {
	counter := 0
	ch := make(chan int)
	mu := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go addnum(ch, &counter, &mu)
	}

	time.Sleep(time.Second)

	fmt.Println("after 10 * 1000 increments, the counter is:", counter)
}
