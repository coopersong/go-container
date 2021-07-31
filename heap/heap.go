package heap

// Sorter is a special slice.
type Sorter interface {
    // Len is the number of elements in the slice.
    Len() int

    // Less reports whether the element with index i
    // must sort before the element with index j.
    //
    // If both Less(i, j) and Less(j, i) are false,
    // then the elements at index i and j are considered equal.
    // Sort may place equal elements in any order in the final result,
    // while Stable preserves the original input order of equal elements.
    //
    // Less must describe a transitive ordering:
    //  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
    //  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
    //
    // Note that floating-point comparison (the < operator on float32 or float64 values)
    // is not a transitive ordering when not-a-number (NaN) values are involved.
    // See Float64Slice.Less for a correct implementation for floating-point values.
    Less(i, j int) bool

    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)

    // Add adds val to the end of slice.
    Add(val interface{})

    // Remove remove the last element of the slice.
    Remove()

    // Set sets val to index i.
    Set(i int, val interface{})

    // Get gets the element with index i.
    Get(i int) interface{}
}

type Heap struct {
    tree Sorter
}

func NewHeap(sorter Sorter) *Heap {
    heap := new(Heap)
    heap.tree = sorter
    return heap
}

// Push pushes val into the heap.
func (h *Heap) Push(val interface{}) {
    h.tree.Add(val)

    n := h.tree.Len()
    index := n - 1
    for index > 0 {
        parent := (index - 1) / 2
        if !h.tree.Less(index, parent) {
            break
        }

        h.tree.Swap(index, parent)

        index = parent
    }
}

func (h *Heap) Pop() interface{} {
    if h.tree.Len() <= 0 {
        return nil
    }

    top := h.tree.Get(0)

    h.tree.Set(0, h.tree.Get(h.tree.Len()-1))
    h.tree.Remove()

    n := h.tree.Len()
    parent := 0
    for parent < n {
        left := parent*2 + 1
        if left >= n {
            break
        }

        right := parent*2 + 2
        index := left
        if right < n && h.tree.Less(right, left) {
            index = right
        }

        if h.tree.Less(parent, index) {
            break
        }

        h.tree.Swap(parent, index)

        parent = index
    }

    return top
}

func (h *Heap) Top() interface{} {
    if h.tree.Len() > 0 {
        return h.tree.Get(0)
    }

    return nil
}

func (h *Heap) Empty() bool {
    return h.tree.Len() == 0
}
