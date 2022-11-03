package main

import "fmt"

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さよなら"
	delete(m, 1)
}

func modSlices(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
}

func main() {
	m := map[int]string{
		1: "1番目",
		2: "2番目",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlices(s)
	fmt.Println(s)
}
