package main

import "os"
func main() {
	args := os.Args
	for _, arg := range args {
		println(arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		println(err.Error())
	} else {
		println(hostname)
	}
}

// go run os.go arg1 "arg 2"