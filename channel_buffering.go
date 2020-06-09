package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "channel"
	messages <- "buffering"
	go func() {messages <- "the third"}()

	fmt.Println(<- messages)
	fmt.Println(<- messages)
	fmt.Println(<- messages)
}