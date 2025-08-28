package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)
	for i := 0; i < r.Len(); i++ {
		r.Value = "value " + fmt.Sprint(i)
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		fmt.Println(p)
	})
}