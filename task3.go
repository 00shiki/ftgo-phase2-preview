package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer close(ch1)
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := range ch1 {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
