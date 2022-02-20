package main

import "fmt"

// another syntax
func main() {
	dataChan := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	for n := range dataChan {
		fmt.Println(n)
	}
}

// func main() {
// 	dataChan := make(chan int, 2)
// 	dataChan <- 789
// 	dataChan <- 343
// 	n := <-dataChan
// 	fmt.Println(n)
// 	n = <-dataChan
// 	fmt.Println(n)
// }
