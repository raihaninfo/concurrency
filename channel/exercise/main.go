package main

import "fmt"

// func sum(a []int, c chan int) {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	c <- sum
// }

// func main() {
// 	a := []int{7, 2}

// 	c := make(chan int)
// 	go sum(a, c)
// 	go sum(a, c)
// 	x, y := <-c, <-c

// 	fmt.Println(x, y)
// }

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
