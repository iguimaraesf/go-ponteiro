package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}
func main() {
	go contador("a")
	go contador("b")
	// sai de modo preemptivo da função
	go func() {
		fmt.Println("função monopolizadora de processamento")
		for {

		}
	}()
	contador("main - sem go routine")
	fmt.Println("main - Hello 1")
	fmt.Println("main - Hello 2")

	time.Sleep(time.Second)
	fmt.Println("main - Acabou!")
}
