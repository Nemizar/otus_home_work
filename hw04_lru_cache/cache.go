package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value any) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok == true {
		c.queue.MoveToFront(el)
		el.Value.(*cacheItem).value = value

		return true
	}

	if c.queue.Len() == c.capacity {
		c.Clear()
	}

	item := &cacheItem{
		key:   key,
		value: value,
	}

	el := c.queue.PushFront(item)
	c.items[item.key] = el

	return false
}

func (c *lruCache) Get(key Key) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	el, ok := c.items[key]
	if ok == false {
		return nil, false
	}

	c.queue.MoveToFront(el)

	return el.Value.(*cacheItem).value, true
}

func (c *lruCache) Clear() {
	if el := c.queue.Back(); el != nil {
		c.queue.Remove(c.queue.Back())

		delete(c.items, el.Value.(*cacheItem).key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
