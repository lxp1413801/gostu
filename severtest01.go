package main

import (
	"fmt"
	"reflect"
	"runtime"
)

var ch chan int = make(chan int, 1)

func main() {
	var a int = 1
	ch <- (a)
	fmt.Println(runtime.NumCPU())
	fmt.Println(reflect.TypeOf(ch))
}
