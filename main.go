package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int, confirmation chan bool) {
	for res := range msg {
		fmt.Printf("a...: Worker: %d Msg: %d\n", workerId, res)
		confirmation <- true
		time.Sleep(time.Second)
	}
}

func main() {
	msg := make(chan int)
	confimation := make(chan bool)

	nworkers := 100
	nrequests := 1000

	for i := 0; i <= nworkers; i++ {
		go worker(i, msg, confimation)
	}

	for i := 0; i <= nrequests; i++ {
		msg <- i
		<-confimation
	}
}
