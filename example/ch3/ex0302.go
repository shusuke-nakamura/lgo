package main

import "fmt"

func main() {
	{
		var data []int // スライスのゼロ値nilで初期化される(nilスライス)
		fmt.Println(data == nil)
	}
	{
		x := []int{} // 空のスライスリテラル
		fmt.Println(x == nil)
	}
}
