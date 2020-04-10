package lru

import "testing"

func TestLRUChan(t *testing.T) {
	lc := NewChan(50)
	for i := 0; i < 100; i++ {
		lc.Put(i)
	}
	lc.Range(func(i interface{}) {
		t.Log("the lru chan range:", i)
	})
}
