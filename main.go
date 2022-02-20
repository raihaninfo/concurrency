package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		// fmt.Println(<-c1)
		// fmt.Println(<-c2)
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

// func main() {
// 	c := make(chan string, 2)
// 	c <- "Hello"
// 	c <- "World"
// 	msg := <-c
// 	fmt.Println(msg)

// 	msg = <-c
// 	fmt.Println(msg)
// }

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
