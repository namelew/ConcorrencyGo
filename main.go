package main

import "fmt"

func main() {
	hello := make(chan string)

	go func() {
		hello <- "hello world"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("a...: default")
	}
}
