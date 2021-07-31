package stack

import (
    "testing"
)

func TestStack(t *testing.T) {
    stack := &Stack{}
    for i := 0; i < 5; i++ {
        stack.Push(i)
    }

    i := 4
    for !stack.Empty() {
        top1 := stack.Top()
        top2 := stack.Pop()
        if top1 != top2 {
            t.Errorf("stack.Top() != stack.Pop()")
            return
        }
        if top1 != i {
            t.Errorf("stack error: i = %d, top = %d", i, top1)
        }
        i--
    }
}
