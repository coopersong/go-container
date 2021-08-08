# go-container

## Getting Started

### stack

```go
package main

import (
    "fmt"
  
    "github.com/coopersong/go-container/stack"
)

func main() {
    // create the stack
    stack := &stack.Stack{}
  
    // push 1-5 to the stack serially
    for i := 1; i <= 5; i++ {
        stack.Push(i)   
    }
    
    // pop all numbers out of the stack
    // pop sequence will be 5, 4, 3, 2, 1
    for !stack.Empty() {
        top := stack.Top()
        fmt.Printf("%d is in the top of the stack\n", top)
        top = stack.Pop()
        fmt.Printf("pop %d out of the stack", top
    }
}
```

### queue

```go
package main

import (
    "fmt"
    
    "github.com/coopersong/go-container/queue"
)

func main() {
    // create the queue
    que := &queue.Queue{}
  
    // push 1-5 to the queue serially
    for i := 1; i <= 5; i++ {
      que.Push(i) 
    }
  
    // pop all numbers out of the queue
    // pop sequence will be 1, 2, 3, 4, 5
    for !que.Empty() {
        front := que.Front()
        fmt.Printf("%d is in the front of the queue\n", front)
        front = que.Pop()
        fmt.Printf("pop %d out of the queue", front)
    }
}
```

### lru

```go
package main

import (
    "fmt"

    "github.com/coopersong/go-container/lru"
)

func main() {
    // create the lru cache capacity 2
    cache, err := lru.NewCache(2)
    if err != nil {
        fmt.Printf("err in lru.NewCache(): %s\n", err.Error())
        return
    }

    cache.Put(1, 1)
    // lru cache: [1]
    cache.Put(2, 2)
    // lru cache: [2, 1]
    a := cache.Get(1)
    // lru cache: [1, 2]
    if a == nil {
        // never come here
        fmt.Printf("cache.Get(1): get no value\n")
        return
    } else {
        fmt.Printf("get %d from lru cache by key 1\n", a)
    }

    b := cache.Get(2)
    // lru cache: [2, 1]
    if b == nil {
        // never come here
        fmt.Printf("cache.Get(2): get no value\n")
        return
    } else {
        fmt.Printf("get %d from lru cache by key 2\n", b)
    }

    sum := a.(int) + b.(int)
    fmt.Printf("%d + %d = %d\n", a, b, sum)

    cache.Put(3, 3)
    // lru cache: [3, 2]
    a = cache.Get(1)
    if a == nil {
        fmt.Printf("cache.Get(1): get no value\n")
    } else {
        // never come here
        fmt.Printf("get %d from lru cache by key 1\n", a)
    }
}
```

## Implementation

### stack

* slice

### queue

* slice

### lru

* map
* two-way linked-list
