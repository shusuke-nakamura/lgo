package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for i := 0; i < 5; i++ {
		fmt.Println("ループ", i)
		fmt.Println(m)
	}
}
