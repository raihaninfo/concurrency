package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"
	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

// func main() {
// 	c := make(chan string)
// 	go count("sheep", c)
// 	for msg := range c {
// 		fmt.Println(msg)
// 	}
// }

// func count(thing string, c chan string) {
// 	for i := 0; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}
// 	close(c)
// }
