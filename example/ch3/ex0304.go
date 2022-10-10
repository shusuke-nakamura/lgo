package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("e:", e)
}
