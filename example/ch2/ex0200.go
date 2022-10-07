// 2章のリスト番号の付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。

package main

import "fmt"

func main() {
	fmt.Println("===== 2.1　基本型 =====")
	fmt.Println("===== 2.1.1 ゼロ値 =====")
	{
		var x int
		var y float32
		var z string

		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Printf("z:|%s|\n", z)
	}

	fmt.Println("===== 2.1.2　リテラル =====")
	fmt.Println("2.1.2.1　整数リテラル")
	fmt.Println(1_234_567) // 1234567

	fmt.Println("2.1.2.2　浮動小数点リテラル")
	fmt.Println(1_234.567_89)

	fmt.Println("2.1.2.3　runeリテラル")
	{
		// A
		fmt.Println("1->\x41")
		fmt.Println("2->\u0041")
		fmt.Println("3->\U00000041")
		// あ
		fmt.Println("1->\u3042")
		fmt.Println("2->\U00003042")
		// 絵文字
		fmt.Println("1->\U0001F600")

		fmt.Println("\n2.1.2.4　文字列リテラル")
		x := `バッククオートを使って"ロー文字列リテラル"を書くことで
改行や二重引用符（ダブルクオート）を文字列中に含めることができる`
		fmt.Println(x)

	}

	fmt.Println("\n2.1.2.5　リテラル型")
	{
		var x float32 = 2 * 0.23 // エラーにならない
		fmt.Println(x)
		var a float64 = 3.14
		var b float64 = 10 / a
		fmt.Println(b) // エラーとならない

		// var c int = "abc" // コンパイル時エラー
		// var s string = 12 // コンパイル時エラー
		var d int
		// d = 12.3 // コンパイル時エラー
		d = 12.0
		fmt.Println("d:", d)
	}

	fmt.Println("===== 2.1.3 論理型 =====")
	{
		var flag bool // 値が代入されないとfalseとなる
		var isAwesom = true
		fmt.Println(flag)
		fmt.Println(isAwesom)
	}

}
