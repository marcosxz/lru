package lru

import (
	"sync"
)

type lRUMap struct {
	sync.RWMutex
	cache map[string]interface{}
	keys  []string
	max   int
}

func NewLRUMap(max int) *lRUMap {
	return &lRUMap{
		max:   max,
		cache: make(map[string]interface{}, max),
		keys:  make([]string, 0, max),
	}
}

func (m *lRUMap) Put(key string, value interface{}) {
	m.Lock()
	if n := len(m.keys); n >= m.max {
		delete(m.cache, m.keys[0])
		m.keys = m.keys[1:n]
	}
	m.cache[key] = value
	m.keys = append(m.keys, key)
	m.Unlock()
}

func (m *lRUMap) Get(key string) interface{} {
	m.RLock()
	value := m.cache[key]
	m.RUnlock()
	return value
}

func (m *lRUMap) Range(fn func(string, interface{})) {
	m.RLock()
	for k, v := range m.cache {
		fn(k, v)
	}
	m.RUnlock()
}
