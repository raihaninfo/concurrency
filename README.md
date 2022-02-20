# concurrency

> concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or at the same time simultaneously partial order,

> Concurrency is when a large task is split into smaller sub-tasks and run at the same time. This post provides an outline of how Go handles concurrency.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

```

> The second one will never be printed here because the first one will never end

> > We use Concurrency to solve this problem

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 0; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

```

## Goroutines

> Goroutines are an integral part of Go. They are the units of execution inside a Go program. The main function is also another goroutine. They all have very small stack sizes, which allows spawning millions of them at the same time. They are extremely lightweight.

https://golangdocs.com/goroutines-in-golang

## Anonymous goroutines

Go has support for anonymous functions. `Goroutines` can be anonymous as well. Here is an example of an anonymous goroutine.

```go
package main

import (
    "fmt"
    "time"
)

func PrintName(f string, l string) {
    fmt.Println(f, l)
}

func main() {
    var i int
    go func() {
        for i = 0; i < 7; i++ {
            fmt.Print(i, " ")
            time.Sleep(100 * time.Millisecond)
        }
    }()
    time.Sleep(1 * time.Second)
    PrintName("John", "Doe")
}
```

## Channel

> Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

> Channel is a communication object using a goroutine

[Channel Article](https://go101.org/article/channel.html)
