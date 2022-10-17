// 3章のリスト番号の付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。
package main

import "fmt"

func main() {
	fmt.Println("===== 3.1　配列 =====")
	{
		var x [3]int
		fmt.Println(x)
	}
	{
		var x = [3]int{10, 20, 30}
		fmt.Println(x)
	}
	{
		var x = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	fmt.Println("----- 配列の比較 -----")
	{
		var x = [...]int{1, 2, 3}
		var y = [3]int{1, 2, 3}
		fmt.Println(x == y)
	}

	fmt.Println("----- 多次元配列のシミュレート -----")
	{
		var x [2][3]int
		fmt.Println(x)
	}

	fmt.Println("----- インデックス（添字）の指定 -----")
	{
		var x [3]int
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])

		var y [2][3]int
		y[0][2] = 12
		y[1][1] = 3
	}

	fmt.Println("----- len -----")
	{
		var x [3]int
		fmt.Println("len(x):", len(x))

		var y [2][3]int
		fmt.Println("len(y):", len(y))
	}

	fmt.Println("===== 3.2　スライス =====")
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])
	}

	{
		var x []int
		fmt.Println(x)
	}

	fmt.Println("----- スライスの比較 -----")
	{
		var x []int
		var y []int
		// fmt.Println(x == y) // エラー
		fmt.Println("x == nil:", x == nil)
		fmt.Println("y !=nil:", y != nil)
	}

	fmt.Println("===== 3.2.1 len =====")
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("len(x):", len(x))
		var y = []int{}
		fmt.Println("len(y):", len(y))
	}

	fmt.Println("----- スライスのスライス -----")
	{
		var x [][]int
		fmt.Println(x)

		var y = [][]int{{0, 1}, {2, 3}, {4, 5}}
		fmt.Println(y)
		fmt.Println(y[1])
		fmt.Println(y[1][1])
	}

	fmt.Println("----- lenの引数 -----")
	{
		var x [][]int
		fmt.Println("len(x):", len(x))
		var y = []int{2, 3, 4, 5}
		fmt.Println("len(y):", len(y))
		var z string = "abc"
		fmt.Println("len(z):", len(z))

		var a = 12
		fmt.Println("a:", a)
		// fmt.Println(len(a)) // エラー
	}

	fmt.Println("===== 3.2.2　append =====")
	{
		var x []int
		fmt.Println(x)
		x = append(x, 10)
		fmt.Println("append(x, 10):", x)
	}

	{
		var x = []int{1, 2, 3}
		x = append(x, 4)
		fmt.Println("append(x, 4):", x)
		x = append(x, 5, 6, 7)
		fmt.Println("append(x, 5, 6, 7):", x)

		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println("x:", x)
		fmt.Println("append(x, y...):", append(x, y...))
		// append(x, y) // エラー
	}

	fmt.Println("===== 3.2.4　make =====")
	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		fmt.Println("x[0], x[4]:", x[0], x[4])
	}
	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		x = append(x, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}
	fmt.Println("----- キャパシティ（第3引数）を指定 -----")
	{
		x := make([]int, 5, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}
	{
		x := make([]int, 0, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}
	{
		x := make([]int, 0, 10)
		x = append(x, 5, 6, 7, 8)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	fmt.Println("===== 3.2.5　スライスの生成方法の選択 =====")
	{
		var data []int
		fmt.Println(data, len(data), cap(data))
		fmt.Println("data == nil:", data == nil)

		var x = []int{}
		fmt.Println(x, len(x), cap(x))
		fmt.Println("x == nil:", x == nil)
	}
	{
		data := []int{2, 4, 6, 8}
		fmt.Println(data, len(data), cap(data))
	}

	fmt.Println("===== 3.3　文字列、rune、バイト =====")
	{
		var s string = "Hello there"
		var b byte = s[6]
		fmt.Println(b)
		fmt.Printf("printfで指定:%c\n", b)
		var c rune = rune(s[6])
		fmt.Println(c)
		fmt.Printf("printfで指定:%c\n", c)
	}

	fmt.Println("----- 文字列に対する「スライス式」 -----")
	{
		var s string = "Hello there"
		fmt.Println("s:", s)
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Printf("s2:%s\ns3:%s\ns4:%s\n", s2, s3, s4)
	}

	fmt.Println("----- 絵文字 -----")
	{
		var s string = "Hello ☀"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println(s)
		fmt.Println(s2)
		fmt.Println(s3)
		fmt.Println(s4)
		fmt.Println("len(s):", len(s))
	}

	fmt.Println("----- 日本語 -----")
	{
		var s string = "こんにちは"
		fmt.Println("s:", s)
		fmt.Println("s[0]:", s[0])
		var b byte = s[6]
		fmt.Println("b:", b)
		var c rune = rune(s[6])
		fmt.Println("c:", c)
	}

	{
		var s string = "こんにちは、みなさん"
		var s2 string = s[4:7]
		var s3 string = s[:5]
		var s4 string = s[6:]
		fmt.Println("s:", s)
		fmt.Println("s2:", s2)
		fmt.Println("s3:", s3)
		fmt.Println("s4:", s4)
	}

	{
		var a rune = 'x'
		var s string = string(a)
		var b byte = 'y'
		var s2 string = string(b)
		fmt.Println("a, s, b, s2:", a, s, b, s2)
	}

	{
		var x int = 65
		var y = string(x)
		fmt.Println("y:", y) // A （65ではない）
	}

	fmt.Println("===== 3.4　マップ =====")
	{
		var nilMap map[string]int
		fmt.Println("nilMap == nil", nilMap == nil)
		fmt.Println("len(nilMap):", len(nilMap))
		fmt.Println("nilMap[\"abc\"]", nilMap["abc"])
		// nilMap["abc"] = 3 // ←パニックになる

		totalWins := map[string]int{}
		fmt.Println("totalWins == nil:", totalWins == nil)
		fmt.Println("totalWins[\"abc\"]", totalWins["abc"])
		totalWins["abc"] = 3
		fmt.Println("totalWins[\"abc\"]", totalWins["abc"])
	}
	{
		totalWins := map[string]int{
			"ゼネターズ":  14,
			"スパローズ":  15,
			"ファルコンズ": 22,
		}
		fmt.Println(totalWins)
	}
	{
		teams := map[string][]string{
			"ライターズ":    []string{"夏目", "森", "国木田"},
			"ナイツ":      []string{"武田", "徳川", "明智"},
			"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
		}
		fmt.Println(teams)
		fmt.Println(teams["ライターズ"])
		fmt.Println(teams["ロビンズ"])

		/*
			teams := map[string][]string {
				"Orcas": []string{"Fred", "Ralph", "Bijou"},
				"Lions": []string{"Sarah", "Peter", "Billie"},
				"Kittens": []string{"Waldo", "Raul", "Ze"},
			}

			fmt.Println(teams) // map[Kittens:[Waldo Raul Ze] Lions:[Sarah Peter Billie] Orcas:[Fred Ralph Bijou]]
			fmt.Println(teams["Lions"]) // [Sarah Peter Billie]
			fmt.Println(teams["Kittens"]) // [Waldo Raul Ze]
		*/

		teams2 := map[string][]string{
			"シャチチーム":  []string{"謙信", "信長", "家康"},
			"ライオンチーム": []string{"レオ", "たか子", "カナ"},
			"猫チーム":    []string{"AKB", "MNB", "FNB"},
		}
		fmt.Println(teams2)
		fmt.Println(teams2["シャチチーム"])
		fmt.Println(teams2["チャチチーム"])
		fmt.Println(teams2["猫チーム"])

		fmt.Println("len(teams2[\"猫チーム\"]):", len(teams2["猫チーム"]))
	}

	{
		ages := make(map[int][]string, 10)
		fmt.Println("ages:", ages)
		fmt.Println("len(ages):", len(ages))
	}

	fmt.Println("===== 3.4.2　カンマOKイディオム =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}
		v, ok := m["hello"]
		fmt.Println(v, ok)

		v, ok = m["world"]
		fmt.Println(v, ok)

		v, ok = m["goodbye"]
		fmt.Println(v, ok)
	}

	fmt.Println("===== 3.4.3 マップからの削除 =====")
	{
		m := map[string]int{
			"hello": 5,
			"world": 0,
		}
		fmt.Println(m)
		v, ok := m["hello"]
		fmt.Println(v, ok)
		delete(m, "hello")
		fmt.Println(m)
		v, ok = m["hello"]
		fmt.Println(v, ok)
		delete(m, "こんにちは")
		fmt.Println(m)
	}
}
