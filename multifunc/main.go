package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generateNumbers(total int) {
	defer wg.Done()
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
		time.Sleep(time.Second)
	}
}

func printNumbers() {
	defer wg.Done()

	for idx := 1; idx <= 10; idx++ {
		fmt.Printf("Printing number %d\n", idx)
		time.Sleep(time.Second)
	}
}

func main() {

	wg.Add(2)
	go printNumbers()
	go generateNumbers(10)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
