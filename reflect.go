package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name string `required:"true" max:"50"`
	Address string `required:"true" max:"50"`
	Email string `required:"true" max:"50"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		fmt.Printf("Field %d: %s %s\n", i, field.Name, field.Type)
		fmt.Println(field.Tag.Get("required"))
		fmt.Println(field.Tag.Get("max"))	
		} 
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if !result {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"example"})
	readField(Person{"John", "bogor", "pH3l9@example.com"})

	person := Person{"adsf", "bogor", "pH3l9@example.com"}
	fmt.Println(isValid(person))
}