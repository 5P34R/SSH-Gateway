package gateway

import "sync"

type CyclicList struct {
	items []string
	index int
	mu    sync.Mutex
}

func NewCyclicList(items []string) *CyclicList {
	return &CyclicList{
		items: items,
	}
}

func (cl *CyclicList) Next() string {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if len(cl.items) == 0 {
		return ""
	}

	item := cl.items[cl.index]
	cl.index = (cl.index + 1) % len(cl.items)
	return item
}
