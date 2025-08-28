// strconv

package main

import (
	"fmt"
	"strconv"
)
func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(boolean)
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}

	number2, err := strconv.Atoi("1000000")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number2)
	}

	binary := strconv.FormatInt(1000000, 2)
	stringInt := strconv.Itoa(1000000)
	fmt.Println(binary)
	fmt.Println(stringInt)
}