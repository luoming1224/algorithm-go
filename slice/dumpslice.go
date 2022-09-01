package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var sl1 = []int{}
	ph1 := (*reflect.SliceHeader)(unsafe.Pointer(&sl1))
	fmt.Printf("empty slice's header is %#v\n", *ph1)
	var sl2 []int
	ph2 := (*reflect.SliceHeader)(unsafe.Pointer(&sl2))
	fmt.Printf("nil slice's header is %#v\n", *ph2)
}
