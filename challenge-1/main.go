package main

import (
	"fmt"
	"sync"
)

var msg string
var waitGroup sync.WaitGroup

func updateMessage(s string) {
	defer waitGroup.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello, world!"

	waitGroup.Add(1)
	go updateMessage("Hello, universe!")
	waitGroup.Wait()
	printMessage()

	waitGroup.Add(1)
	go updateMessage("Hello, cosmos!")
	waitGroup.Wait()
	printMessage()

	waitGroup.Add(1)
	go updateMessage("Hello, go!")
	waitGroup.Wait()
	printMessage()
}
