package main

import "fmt"

func main() {
	fmt.Println("===== 7.2.7　iotaと列挙型 =====")
	{
		type MailCategory int // メールの分類
		const (
			Uncategorized  MailCategory = iota // 未分類
			Personal                           // 個人的
			Spam                               // 迷惑メール
			Social                             // ソーシャル
			Advertisements                     // 広告
		)

		m := Personal
		fmt.Println("Personal:", m)
		m = Uncategorized
		fmt.Println("Uncategorized", m)
	}

	fmt.Println("----- ビットフィールド -----")
	{
		type BitField int
		const (
			Field1 BitField = 1 << iota
			Field2
			Field3
			Field4
		)
		fmt.Println("Field1:", Field1)
		fmt.Println("Field2:", Field2)
		fmt.Println("Field3:", Field3)
		fmt.Println("Field4:", Field4)
	}

	fmt.Println("----- 「_」 -----")
	{
		type SomeValue int
		const (
			_ SomeValue = iota
			Value1
			Value2
			Value3
			Value4
		)
		fmt.Println("Value1:", Value1)
		fmt.Println("Value2:", Value2)
		fmt.Println("Value3:", Value3)
		fmt.Println("Value4:", Value4)
	}

	fmt.Println("----- 無効を表す定数 -----")
	{
		type SomeValue int
		const (
			Invalid SomeValue = iota
			Value1
			Value2
			Value3
			Value4
		)
		fmt.Println("Invalid:", Invalid)
		fmt.Println("Value1", Value1)
		fmt.Println("Value2", Value2)
		fmt.Println("Value3", Value3)
		fmt.Println("Value4", Value4)
	}
}
