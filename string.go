package main

import (
	"fmt"
	"strings"
)

func main()  {
	// contains, split, tolower, toupper, replace, trim
	fmt.Println(strings.Contains("hello world", "world"))
	fmt.Println(strings.Split("hello world", " "))
	fmt.Println(strings.ToLower("HELLO WORLD"))
	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.Replace("hello world", "world", "planet", 1))
	fmt.Println(strings.Trim("    hello world    ", " "))
}