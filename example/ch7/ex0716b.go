package main

import "fmt"

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i4 := i.(int)
	fmt.Println(i4)
}
