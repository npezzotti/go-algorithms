package lru_cache

import "testing"

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(2)

	lruCache.Put(1, 1)
	lruCache.Put(2, 2)

	if num := lruCache.Get(1); num != 1 {
		t.Errorf("got %d, expected %d", num, 1)
	}

	lruCache.Put(3, 3)

	if num := lruCache.Get(2); num != -1 {
		t.Errorf("got %d, expected %d", num, -1)
	}

	lruCache.Put(4, 4)

	if num := lruCache.Get(1); num != -1 {
		t.Errorf("got %d, expected %d", num, -1)
	}
	if num := lruCache.Get(3); num != 3 {
		t.Errorf("got %d, expected %d", num, 3)
	}
	if num := lruCache.Get(4); num != 4 {
		t.Errorf("got %d, expected %d", num, 4)
	}
}
