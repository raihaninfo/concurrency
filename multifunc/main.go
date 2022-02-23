package main

import (
	"fmt"
	"sync"
	"time"
)

func generateNumbers(total int, wg *sync.WaitGroup) {
	defer wg.Done()
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
		time.Sleep(time.Second)
	}
}

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()

	for idx := 1; idx <= 10; idx++ {
		fmt.Printf("Printing number %d\n", idx)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go printNumbers(&wg)
	go generateNumbers(10, &wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
