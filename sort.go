package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}
func main() {
	users := UserSlice{
		{"Dwi", 23},
		{"Muhammad", 20},
		{"Rudi", 25},
		{"Susanto", 28},
	}
	sort.Sort(users)
	fmt.Print(users)
}