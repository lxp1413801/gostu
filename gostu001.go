package main

import (
	"fmt"
	"time"
)

var ch1 chan int = make(chan int, 3)

func task1() {
	var x int = 0
	fmt.Println("call task1")
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("task1 send message:%d\r\n", x)
		ch1 <- (x)
		ch1 <- (x + 1)
		ch1 <- (x + 2)
		x += 3
	}
}

func task2() {
	var x int
	for {
		x = <-ch1
		fmt.Printf("task2 received msg:%d\r\n", x)
		x = <-ch1
		fmt.Printf("task2 received msg:%d\r\n", x)
		x = <-ch1
		fmt.Printf("task2 received msg:%d\r\n\r\n", x)
	}
}

func main() {
	fmt.Println("hello go!!!世界的第一个golong程序")
	go task1()
	go task2()
	for {
		time.Sleep(1000 * time.Millisecond)
	}

}
