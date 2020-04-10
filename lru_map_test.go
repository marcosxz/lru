package lru

import (
	"testing"
)

func TestLRUMap(t *testing.T) {
	lm := NewMap(50)
	for i := 0; i < 100; i++ {
		lm.Put(i, i)
	}
	for i := 0; i < 100; i++ {
		t.Logf("lru map key%d:%v", i, lm.Get(i))
	}
	lm.Range(func(s, i interface{}) {
		t.Log("lru map range: key =", s, "; value =", i)
	})
}
