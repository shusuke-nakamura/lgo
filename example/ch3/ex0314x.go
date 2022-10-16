package main

import (
	"fmt"
	"reflect"
)

// reflect.TypeOf(x) で xの型がわかります（14章参照）
func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println("x:", x, reflect.TypeOf(x))
	d := [4]int{5, 6, 7, 8}
	fmt.Println("d:", d, reflect.TypeOf(d))
	y := make([]int, 2)
	fmt.Println("y:", y, reflect.TypeOf(y))
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println("d:", d, reflect.TypeOf(d))
	copy(d[2:3], x)
	fmt.Println("d:", d)

	copy(d[2:4], x)
	fmt.Println("d:", d)

	copy(d[2:], x)
	fmt.Println("d:", d)

	copy(d[1:], x)
	fmt.Println("d:", d)
}
