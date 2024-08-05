package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan rune, 3)
	ch2 := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer close(ch1)
		for i := 'a'; i <= 'j'; i++ {
			ch1 <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer close(ch2)
		for i := 1; i <= 10; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := range ch1 {
			fmt.Println(string(i))
		}
	}()

	go func() {
		defer wg.Done()
		for i := range ch2 {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}
