package main

import "fmt"

// 型Personを定義
type Person struct {
	LastName  string
	FirstName string
	Age       int
}

// 型Personに付随するメソッドStringを定義(PersonにメソッドStringを付加)
func (p Person) String() string {
	return fmt.Sprintf("%s %s:年齢%d歳", p.LastName, p.FirstName, p.Age)
}

func main() {
	p := Person{
		LastName:  "武田",
		FirstName: "信玄",
		Age:       52,
	}

	output := p.String()
	fmt.Println(output)
}
