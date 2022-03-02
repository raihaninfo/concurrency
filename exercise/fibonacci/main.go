// fibonacci algorithm
package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	go worker(jobs, result)
	go worker(jobs, result)
	go worker(jobs, result)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j < 100; j++ {
		fmt.Println(<-result)
	}
}

func worker(jobs <-chan int, result chan<- int) {
	for n := range jobs {
		result <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
