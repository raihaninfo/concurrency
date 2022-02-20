package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
	time.Sleep(time.Second * 2)
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
