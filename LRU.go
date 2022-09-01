package main

import (
	"container/list"
)

var LRUMap = make(map[string]*list.Element)
var LRUList *list.List = list.New()

func insertLRU(key string) {
	if object, exist := LRUMap[key]; exist {
		LRUList.MoveToFront(object)
	} else {
		object = new(list.Element)
		object.Value = key
		LRUList.PushFront(object)
	}
}

func test() {

}
