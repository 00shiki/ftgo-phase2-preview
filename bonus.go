package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	errs := make(chan error, 2)
	quit := make(chan int)

	go func() {
		defer close(ch1)
		defer close(ch2)
		defer close(errs)
		for i := 1; i <= 22; i++ {
			if i > 20 {
				errs <- fmt.Errorf("number %d is greated than 20", i)
			} else if i%2 == 1 {
				ch1 <- i
			} else {
				ch2 <- i
			}
		}
		quit <- 0
	}()

	for {
		select {
		case num := <-ch1:
			fmt.Printf("Received an odd number: %d\n", num)
		case num := <-ch2:
			fmt.Printf("Received an even number: %d\n", num)
		case err := <-errs:
			fmt.Printf("Error: %v\n", err)
		case <-quit:
			return
		}
	}
}
