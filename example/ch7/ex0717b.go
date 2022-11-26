package main

import (
	"errors"
	"fmt"
	"os"
)

type MyInt int

func main() {
	i, err := do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(i)
}

func do() (MyInt, error) {
	var i any
	var mine MyInt = 20
	i = mine
	i4, ok := i.(int)
	if !ok {
		mes := fmt.Sprintf("%vの値が想定外です。", i)
		return 0, errors.New(mes)
	}
	return MyInt(i4), nil
}
