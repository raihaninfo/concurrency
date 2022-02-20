package main

import "fmt"

func main() {
	dataChan := make(chan int, 2)

	dataChan <- 789
	dataChan <- 343

	n := <-dataChan
	fmt.Println(n)

	n = <-dataChan
	fmt.Println(n)

}
