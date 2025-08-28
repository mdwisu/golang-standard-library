package main

import (
	"container/list"
	"fmt"
)

func main() { 
	var data *list.List = list.New()
	data.PushBack("Muhammad")
	data.PushBack("Dwi")
	data.PushBack("Susanto")

	var head *list.Element = data.Front()

	next := head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}