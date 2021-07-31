package lru

import (
    "errors"
)

type Cache struct {
    capacity int
    head     *node
    tail     *node
    length   int
    mp       map[interface{}]*node
}

type node struct {
    left  *node
    right *node
    key   interface{}
    val   interface{}
}

var (
    ErrInvalidCapacity = errors.New("invalid capacity")
)

func NewCache(capacity int) (*Cache, error) {
    if capacity <= 0 {
        return nil, ErrInvalidCapacity
    }

    cache := new(Cache)
    cache.capacity = capacity
    cache.mp = make(map[interface{}]*node)

    return cache, nil
}

func (c *Cache) Put(key, val interface{}) {
    if node, ok := c.mp[key]; ok {
        node.val = val
        c.Get(key)
        return
    }

    if c.length == 0 {
        node := c.registerNode(key, val)
        c.head = node
        c.tail = node
        c.length++
        return
    }

    if c.length >= c.capacity {
        c.unregisterNode(c.tail)
        c.length--
        if c.tail.left == nil {
            c.head = nil
            c.tail = nil
            node := c.registerNode(key, val)
            c.head = node
            c.tail = node
            c.length++
            return
        }
        c.tail = c.tail.left
        c.tail.right = nil
    }

    node := c.registerNode(key, val)
    c.head.left = node
    node.right = c.head
    c.head = node
    c.length++
}

func (c *Cache) Get(key interface{}) interface{} {
    node, ok := c.mp[key]
    if !ok {
        return nil
    }

    if node.left == nil {
        // No need to adjust.
        return node.val
    }

    node.left.right = node.right
    if node.right == nil {
        c.tail = node.left
    } else {
        node.right.left = node.left
    }

    node.left = nil
    node.right = c.head
    c.head.left = node
    c.head = node

    return node.val
}

func (c *Cache) registerNode(key, val interface{}) *node {
    node := new(node)

    node.key = key
    node.val = val

    c.mp[key] = node

    return node
}

func (c *Cache) unregisterNode(node *node) {
    delete(c.mp, node.key)
}
