package deque

import (
    "errors"
)

type Deque struct {
    head *node
    tail *node
}

type node struct {
    val   interface{}
    left  *node
    right *node
}

var (
    ErrNoElement = errors.New("no element")
)

func NewDeque() *Deque {
    return new(Deque)
}

func (d *Deque) PushBack(v interface{}) {
    if d.tail == nil {
        node := &node{
            val: v,
        }
        d.head = node
        d.tail = node
        return
    }

    node := &node{
        val: v,
    }
    d.tail.right = node
    node.left = d.tail
    d.tail = node
}

func (d *Deque) PushFront(v interface{}) {
    if d.head == nil {
        node := &node{
            val: v,
        }
        d.head = node
        d.tail = node
        return
    }

    node := &node{
        val: v,
    }
    d.head.left = node
    node.right = d.head
    d.head = node
}

func (d *Deque) Empty() bool {
    return d.head == nil
}

func (d *Deque) Front() interface{} {
    if d.head == nil {
        return nil
    }

    return d.head.val
}

func (d *Deque) Back() interface{} {
    if d.tail == nil {
        return nil
    }

    return d.tail.val
}

func (d *Deque) PopFront() error {
    if d.Empty() {
        return ErrNoElement
    }

    if d.tail == d.head {
        d.tail = nil
    }

    d.head = d.head.right
    if d.head != nil {
        d.head.left = nil
    }

    return nil
}

func (d *Deque) PopBack() error {
    if d.Empty() {
        return ErrNoElement
    }

    if d.head == d.tail {
        d.head = nil
    }

    d.tail = d.tail.left
    if d.tail != nil {
        d.tail.right = nil
    }

    return nil
}
