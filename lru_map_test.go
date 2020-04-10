package lru

import (
	"fmt"
	"testing"
)

func TestLRUMap(t *testing.T) {
	lm := NewLRUMap(50)
	for i := 0; i < 100; i++ {
		lm.Put(fmt.Sprintf("key%d", i), i)
	}
	for i := 0; i < 100; i++ {
		t.Logf("lru map key%d:%v", i, lm.Get(fmt.Sprintf("key%d", i)))
	}
	lm.Range(func(s string, i interface{}) {
		t.Log("lru map range: key =", s, "; value =", i)
	})
}
