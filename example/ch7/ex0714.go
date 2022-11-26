package main

import "fmt"

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // 型アサーション。iをMyInt型と仮定(assert)して値(20)を貰う。
	fmt.Println(i2)
}
