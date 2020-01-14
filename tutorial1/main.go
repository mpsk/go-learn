package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)

	fmt.Println(c)

	fmt.Println("=> pinger")
	go pinger(c)
	fmt.Println("=> ponger")
	go ponger(c)
	fmt.Println("=> printer")
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
