package main

import "fmt"

type Score int
type HighScore Score

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

type Employee Person

func (s Score) Double() Score {
	return s * 2
}

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(s, hs)
	hss := hs + 20
	fmt.Println(hss)

	s = 200
	hs = 300
	fmt.Println(s.Double())
	fmt.Println(Score(hs).Double())
}
