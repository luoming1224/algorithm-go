package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	key string		// Node中存储val外，也需要存储Key，否则淘汰时，无法从map中删除对应项
	val string
}

type LRUCache struct {
	capacity int
	lruMap map[string]*list.Element
	lruList *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	lruCache := &LRUCache{
		capacity: capacity,
		lruMap: make(map[string]*list.Element),
		lruList: list.New(),
	}
	return lruCache
}

func (cache *LRUCache)Put(key, val string) {
	if object, exist := cache.lruMap[key]; exist {
		object.Value = &Node{				// 如果存在，更新值
			key: key,
			val: val,
		}
		cache.lruList.MoveToFront(object)
	} else {
		if len(cache.lruMap) >= cache.capacity {
			o := cache.lruList.Back()
			if v, ok := o.Value.(*Node); ok {
				delete(cache.lruMap, v.key)
				cache.lruList.Remove(o)
			}
		}

		node := &Node{
			key: key,
			val: val,
		}
		e := cache.lruList.PushFront(node)	// 注意：往list存Node即可，list会封装Element对象并返回
		cache.lruMap[key] = e				// map中存入插入list时返回的element
	}
}

func (cache *LRUCache)Get(key string) string {
	if o, exist := cache.lruMap[key]; exist {
		cache.lruList.MoveToFront(o)
		if v, ok := o.Value.(*Node); ok {
			return v.val
		}
	}
	return ""
}

func (cache *LRUCache)PrintElement()  {
	fmt.Println("==========")
	cacheLen := cache.lruList.Len()
	o := cache.lruList.Front()
	for i := 0; i < cacheLen; i++ {
		if v, ok := o.Value.(*Node); ok {
			fmt.Println("key:", v.key, "val:", v.val)
		}
		o = o.Next()
	}
}

func main() {
	lruCache := NewLRUCache(3)
	lruCache.Put("a", "1")
	lruCache.Put("b", "2")
	lruCache.Put("c", "3")
	lruCache.PrintElement()
	lruCache.Put("d", "4")
	lruCache.PrintElement()
	lruCache.Get("b")
	lruCache.PrintElement()
	lruCache.Put("c", "5")
	lruCache.PrintElement()
}
