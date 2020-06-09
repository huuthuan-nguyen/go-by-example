package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		messages <- "ping"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		messages <- "pong"
	}()

	ping := <- messages

	pong := <- messages

	fmt.Println(ping)

	fmt.Println(pong)
	

}