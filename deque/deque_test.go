package deque

import (
    "testing"
)

func TestDeque_1(t *testing.T) {
    deque := NewDeque()
    for i := 0; i <= 5; i++ {
        deque.PushBack(i)
    }

    i := 0
    for !deque.Empty() {
        front := deque.Front()
        deque.PopFront()
        if front != i {
            t.Errorf("front should be %d but %d", i, front)
            return
        }
        i++
    }
}

func TestDeque_2(t *testing.T) {
    deque := NewDeque()
    for i := 0; i <= 5; i++ {
        deque.PushFront(i)
    }

    i := 5
    for !deque.Empty() {
        front := deque.Front()
        deque.PopFront()
        if front != i {
            t.Errorf("front should be %d but %d", i, front)
            return
        }
        i--
    }
}

func TestDeque_3(t *testing.T) {
    deque := NewDeque()
    for i := 0; i <= 5; i++ {
        deque.PushBack(i)
    }

    i := 5
    for !deque.Empty() {
        back := deque.Back()
        deque.PopBack()
        if back != i {
            t.Errorf("back should be %d but %d", i, back)
            return
        }
        i--
    }
}

func TestDeque_4(t *testing.T) {
    deque := NewDeque()
    for i := 0; i <= 5; i++ {
        deque.PushFront(i)
    }

    i := 0
    for !deque.Empty() {
        back := deque.Back()
        deque.PopBack()
        if back != i {
            t.Errorf("back should be %d but %d", i, back)
            return
        }
        i++
    }
}

func TestDeque_5(t *testing.T) {
    deque := NewDeque()
    for i := 0; i <= 6; i++ {
        if i%2 == 0 {
            deque.PushFront(i)
        } else {
            deque.PushBack(i)
        }
    }

    i := 6
    for !deque.Empty() {
        front := deque.Front()
        back := deque.Back()
        if i%2 == 0 {
            deque.PopFront()
            if front != i {
                t.Errorf("front should be %d but %d", i, front)
                return
            }
        } else {
            deque.PopBack()
            if back != i {
                t.Errorf("back should be %d but %d", i, back)
                return
            }
        }
        i--
    }
}
