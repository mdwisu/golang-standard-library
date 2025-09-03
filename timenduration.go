package main

import (
	"fmt"
	"time"
)

func main() {
	// time now, date, parse
	now := time.Now()
	fmt.Println(now.Local())
	fmt.Println(now.UTC())
	fmt.Println(now.Year())

	date := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(date)
	fmt.Println(date.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("error", err.Error())
	}
	fmt.Println(valueTime)
	fmt.Println(valueTime.Year())

	// duration
	var d time.Duration = time.Hour*24 + time.Minute*30
	fmt.Println(d.Hours())
	fmt.Println(d.Minutes())
	fmt.Println(d.Seconds())
	fmt.Println(d.Milliseconds())
	fmt.Println(d.Nanoseconds())
}