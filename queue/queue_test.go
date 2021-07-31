package queue

import (
    "testing"
)

func TestQueue(t *testing.T) {
    que := &Queue{}
    for i := 0; i < 5; i++ {
        que.Push(i)
    }

    i := 0
    for !que.Empty() {
        front1 := que.Front()
        front2 := que.Pop()
        if front1 != front2 {
            t.Errorf("que.Front() != que.Pop()")
            return
        }
        if front1 != i {
            t.Errorf("queue error: i = %d, front = %d", i, front1)
            return
        }
        i++
    }
}
