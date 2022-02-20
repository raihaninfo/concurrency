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

