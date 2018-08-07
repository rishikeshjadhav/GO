package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Welcome to Channels")
	ch := make(chan int, 10)
	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		i := 42
		ch <- i
		wg.Done()
	}()

	wg.Wait()
}
