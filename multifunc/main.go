package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func generateNumbers(total int, ch chan<- int) {
	defer wg.Done()
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("sending %d to channel\n", idx)
		ch <- idx
	}
}

func printNumbers(ch <-chan int) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("read %d from channel\n", num)
	}
}

func main() {
	numberChan := make(chan int)
	wg.Add(2)
	go printNumbers(numberChan)

	generateNumbers(3, numberChan)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
