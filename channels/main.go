package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go func(ch <-chan int) {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			ch <- 42
			wg.Done()
		}(ch)
	}
	wg.Wait()

	ch1 := make(chan int, 50)
	wg.Add(2)
	go func(ch1 <-chan int) {
		for i := range ch1 {
			fmt.Println(i)
		}
		wg.Done()
	}(ch1)
	go func(ch1 chan<- int) {
		ch1 <- 421
		ch1 <- 200
		close(ch1)
		wg.Done()
	}(ch1)
	wg.Wait()
}
