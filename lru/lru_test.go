package lru

import (
    "testing"
)

type operation struct {
    params []int
    expect int
}

func TestNewCache(t *testing.T) {
    _, err := NewCache(-2)
    if err == nil {
        t.Error("no error with invalid capacity")
        return
    }
}

func TestLRU_1(t *testing.T) {
    cache, err := NewCache(2)
    if err != nil {
        t.Errorf("NewCache(): %s", err.Error())
        return
    }

    ops := []operation{
        {params: []int{1, 1}},
        {params: []int{2, 2}},
        {params: []int{1}, expect: 1},
        {params: []int{3, 3}},
        {params: []int{2}, expect: -1},
        {params: []int{4, 4}},
        {params: []int{1}, expect: -1},
        {params: []int{3}, expect: 3},
        {params: []int{4}, expect: 4},
    }

    for _, op := range ops {
        if len(op.params) == 2 {
            // Put
            cache.Put(op.params[0], op.params[1])
        } else if len(op.params) == 1 {
            // Get
            val := cache.Get(op.params[0])
            if val == nil {
                if op.expect != -1 {
                    t.Errorf("get nil but expect %d", op.expect)
                    return
                }
            } else {
                if val != op.expect {
                    t.Errorf("get %v but expect %v", val, op.expect)
                    return
                }
            }
        }
    }
}

func TestLRU_2(t *testing.T) {
    cache, err := NewCache(2)
    if err != nil {
        t.Errorf("NewCache(): %s", err.Error())
        return
    }

    ops := []operation{
        {params: []int{2}, expect: -1},
        {params: []int{2, 6}},
        {params: []int{1}, expect: -1},
        {params: []int{1, 5}},
        {params: []int{1, 2}},
        {params: []int{1}, expect: 2},
        {params: []int{2}, expect: 6},
    }

    for _, op := range ops {
        if len(op.params) == 2 {
            // Put
            cache.Put(op.params[0], op.params[1])
        } else if len(op.params) == 1 {
            // Get
            val := cache.Get(op.params[0])
            if val == nil {
                if op.expect != -1 {
                    t.Errorf("get nil but expect %d", op.expect)
                    return
                }
            } else {
                if val != op.expect {
                    t.Errorf("get %v but expect %v", val, op.expect)
                    return
                }
            }
        }
    }
}
