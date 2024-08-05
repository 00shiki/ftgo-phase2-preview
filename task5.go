package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	quit := make(chan int)

	go func() {
		defer close(ch1)
		defer close(ch2)
		for i := 1; i <= 20; i++ {
			if i%2 == 1 {
				ch1 <- i
			} else {
				ch2 <- i
			}
		}
		quit <- 0
	}()

	for i := 0; i < 20; i++ {
		select {
		case num := <-ch1:
			fmt.Printf("Received an odd number: %d\n", num)
		case num := <-ch2:
			fmt.Printf("Received an even number: %d\n", num)
		case <-quit:
			return
		}
	}
}
