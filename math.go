// math
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Math Pi", 3.14)
	fmt.Println("Math Ceil", math.Ceil(3.14))
	fmt.Println("Math Floor", math.Floor(3.14))
	fmt.Println("Math Round", math.Round(3.14))
	fmt.Println("Math Max", math.Max(10, 20))
	fmt.Println("Math Min", math.Min(10, 20))
	fmt.Println("Math Sqrt", math.Sqrt(16))
	fmt.Println("Math Pow", math.Pow(4, 2))
	fmt.Println("Math Mod", 10 % 3)
}