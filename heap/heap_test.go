package heap

import (
    "sort"
    "testing"
)

type Slice []int

func (s *Slice) Add(val interface{}) {
    *s = append(*s, val.(int))
}

func (s *Slice) Less(i, j int) bool {
    return (*s)[i] < (*s)[j]
}

func (s *Slice) Len() int {
    return len(*s)
}

func (s *Slice) Swap(i, j int) {
    (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Slice) Remove() {
    *s = (*s)[:len(*s)-1]
}

func (s *Slice) Get(i int) interface{} {
    return (*s)[i]
}

func (s *Slice) Set(i int, val interface{}) {
    (*s)[i] = val.(int)
}

func TestHeap(t *testing.T) {
    var slice Slice
    heap := NewHeap(&slice)

    var sortedSlice Slice

    for i := 16; i <= 20; i++ {
        if i%2 == 0 {
            heap.Push(i - 15)
            sortedSlice = append(sortedSlice, i-15)
        } else {
            heap.Push(i + 15)
            sortedSlice = append(sortedSlice, i+15)
        }
    }

    sort.Sort(&sortedSlice)
    index := 0
    for !heap.Empty() {
        top1 := heap.Top()
        top2 := heap.Pop()
        if top1 != top2 {
            t.Errorf("heap.Top() != heap.Pop()")
            return
        }
        if top1 != sortedSlice[index] {
            t.Errorf("heap error: index = %d, top = %d", index, top1)
            return
        }
        index++
    }
}
