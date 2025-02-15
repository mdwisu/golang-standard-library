package main

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "dwi" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := GetById("sdf")
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("validation error")
		} else if errors.Is(err, NotFoundError) {
			println("not found error")
		} else {
			println("unknown error")
		}
	}
}