package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

// another syntax
func main() {
	dataChan := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}
		wg.Wait()
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
