# concurrency

concurrency is the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or at the same time simultaneously partial order,

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
