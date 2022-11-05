package main

import "fmt"

func main() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName: "Pat",
		// MiddleName: "Perry", // ←コンパイルエラー
		LastName: "Peterson",
	}
	fmt.Println(p)
}
