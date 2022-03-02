package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{7, 2}

	c := make(chan int)
	go sum(a, c)
	go sum(a, c)
	x, y := <-c, <-c

	fmt.Println(x, y)
}
