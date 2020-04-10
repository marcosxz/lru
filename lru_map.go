package lru

import (
	"sync"
)

type Map struct {
	sync.RWMutex
	cache map[interface{}]interface{}
	keys  []interface{}
	max   int
}

func NewMap(max int) *Map {
	return &Map{
		max:   max,
		cache: make(map[interface{}]interface{}, max),
		keys:  make([]interface{}, 0, max),
	}
}

func (m *Map) Put(key, value interface{}) {
	m.Lock()
	if n := len(m.keys); n >= m.max {
		delete(m.cache, m.keys[0])
		m.keys = m.keys[1:n]
	}
	m.cache[key] = value
	m.keys = append(m.keys, key)
	m.Unlock()
}

func (m *Map) Get(key interface{}) interface{} {
	m.RLock()
	value := m.cache[key]
	m.RUnlock()
	return value
}

func (m *Map) Range(fn func(interface{}, interface{})) {
	m.RLock()
	for k, v := range m.cache {
		fn(k, v)
	}
	m.RUnlock()
}

func (m *Map) Len() int {
	m.RLock()
	n := len(m.keys)
	m.RUnlock()
	return n
}
