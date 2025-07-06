package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "первая горутина"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "первая горутина"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch3 <- "первая горутина"
	}()

	select { //выводит первую исполненную рутину
	case result := <-ch1:
		fmt.Println(result)
	case result := <-ch2:
		fmt.Println(result)
	case result := <-ch3:
		fmt.Println(result)
	}

}
