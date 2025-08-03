package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(s)
}

func main() {
	var waitGroup sync.WaitGroup

	words := []string{
		"alpha",
		"bravo",
		"charlie",
		"delta",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
		"india",
	}

	waitGroup.Add(len(words))

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d:%s", i, word), &waitGroup)
	}

	waitGroup.Wait()
}
