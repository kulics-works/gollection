package gollection

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	var hashmap = HashMapOf[string, int]()
	if hashmap.Count() != 0 {
		t.Fatal("map count not eq 0")
	}
	hashmap.Put("111", 1)
	if hashmap.Count() != 1 {
		t.Fatal("map count not eq 1")
	}
	hashmap.Put("111", 2)
	if hashmap.Count() != 1 {
		t.Fatal("map count not eq 1")
	}
	if v, ok := hashmap.TryGet("111").Get(); !ok || v != 2 {
		t.Fatal("map value not eq 2")
	}
	var strkey = fmt.Sprintf("%d", 111)
	hashmap.Put(strkey, 3)
	if hashmap.Count() != 1 {
		t.Fatal("map count not eq 1")
	}
	if v, ok := hashmap.TryGet(strkey).Get(); !ok || v != 3 {
		t.Fatal("map value not eq 3")
	}
	var hashmap2 = HashMapOf[int, int]()
	if hashmap2.Count() != 0 {
		t.Fatal("map count not eq 0")
	}
	hashmap2.Put(111, 1)
	if hashmap2.Count() != 1 {
		t.Fatal("map count not eq 1")
	}
	hashmap2.Put(111, 2)
	if hashmap2.Count() != 1 {
		t.Fatal("map count not eq 1")
	}
	if v, ok := hashmap2.TryGet(111).Get(); !ok || v != 2 {
		t.Fatal("map value not eq 2")
	}
	var _ AnyMap[string, int] = hashmap
	var _ AnyMutableMap[string, int] = hashmap
}
