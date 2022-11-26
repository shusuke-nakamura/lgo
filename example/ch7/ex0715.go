package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(string)
	fmt.Println(i2)
}
