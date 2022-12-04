package main

import "fmt"

type BitField int

const (
	Field1 BitField = 1 << iota
	Field2
	Field3
	Field4
)

type SomeCategory int

const (
	_ SomeCategory = iota
	CategoryX1
	CategoryX2
	CategoryX3
)

type SomeCategory2 int

const (
	CategoryY1 SomeCategory2 = iota + 3
	CategoryY2
	CategoryY3
)

func main() {
	fmt.Println("=== 7.2.7 iotaと列挙型 ===")

	fmt.Println(Field1)
	fmt.Println(Field2)
	fmt.Println(Field3)
	fmt.Println(Field4)

	fmt.Println("---")
	fmt.Println(CategoryX1)
	fmt.Println(CategoryX2)

	fmt.Println("---")
	fmt.Println(CategoryY1)
	fmt.Println(CategoryY2)
}
