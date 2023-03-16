package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello World"
	}()

	result := <-hello
	fmt.Println(result)

	queue := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			// Só executa novamente quando o valor for lido.
			queue <- i
			i++
		}
	}()

	for x := range queue {
		fmt.Println(x)
	}
}
