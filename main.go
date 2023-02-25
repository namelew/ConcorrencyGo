package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("a: Começou...")

	go func() {
		for {

		}
	}()

	time.Sleep(time.Second)
	fmt.Println("a: Terminou")
}
