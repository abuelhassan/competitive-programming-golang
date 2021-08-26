package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type cacheElement struct {
	key       int
	value     int
	frequency int
}

type LFUCache struct {
	capacity  int
	dll       *list.List
	pointers  map[int]*list.Element
	frequency map[int]*list.Element // maps frequency to most recent used node with that frequency.
}

func newCache(capacity int) LFUCache {
	return LFUCache{
		capacity:  capacity,
		dll:       list.New(),
		pointers:  make(map[int]*list.Element, capacity),
		frequency: make(map[int]*list.Element, capacity),
	}
}

func (lfu *LFUCache) Get(key int) int {
	ptr, ok := lfu.pointers[key]
	if !ok { // key not cached
		return -1
	}
	lfu.incrementFrequency(key)
	return ptr.Value.(cacheElement).value
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}

	// Update if already exists.
	if ptr, ok := lfu.pointers[key]; ok {
		node := ptr.Value.(cacheElement)
		node.value = value
		ptr.Value = node
		lfu.incrementFrequency(key)
		return
	}

	// Invalidate one element if needed.
	if lfu.dll.Len() == lfu.capacity {
		// Invalidate LFU and in case of ties choose LRU.
		ptr := lfu.dll.Front()
		node := ptr.Value.(cacheElement)

		lfu.dll.Remove(ptr)
		delete(lfu.pointers, node.key)
		if lfu.frequency[node.frequency] == ptr {
			delete(lfu.frequency, node.frequency)
		}
	}

	// Insert new element.
	ptr := lfu.dll.PushFront(cacheElement{key: key, value: value, frequency: 0})
	lfu.pointers[key] = ptr
	lfu.frequency[0] = ptr
	lfu.incrementFrequency(key)
}

// incrementFrequency assumes that the element already exist.
func (lfu *LFUCache) incrementFrequency(key int) {
	ptr := lfu.pointers[key]
	node := ptr.Value.(cacheElement)

	// Move element to latest in current frequency.
	if old := lfu.frequency[node.frequency]; old != ptr {
		lfu.dll.MoveAfter(ptr, old)
	}

	// Update value.
	node.frequency++
	ptr.Value = node

	// Update current frequency.
	if ptr.Prev() != nil && ptr.Prev().Value.(cacheElement).frequency == node.frequency-1 {
		lfu.frequency[node.frequency-1] = ptr.Prev()
	} else {
		delete(lfu.frequency, node.frequency-1)
	}

	// Update new frequency.
	if old, ok := lfu.frequency[node.frequency]; ok {
		lfu.dll.MoveAfter(ptr, old)
	}
	lfu.frequency[node.frequency] = ptr
}

func main() {
	// TODO: add unit tests
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	cache := newCache(2)
	cache.Put(3, 1)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Put(4, 4)
	fmt.Println(cache.Get(2))

	var n int
	_, _ = fmt.Fscan(r, &n)
	_, _ = fmt.Fprintf(w, "input: %d", n)
}
