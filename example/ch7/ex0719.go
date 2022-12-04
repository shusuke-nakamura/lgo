package main

import (
	"errors"
	"fmt"
)

type treeNode struct {
	val    treeVal
	lchild *treeNode
	rchild *treeNode
}

type treeVal interface {
	isToken()
}

type nunmber int

func (nunmber) isToken() {}

type operator func(int, int) int

func (operator) isToken() {}

func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

var operators = map[string]operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

func walkTree(t *treeNode) (int, error) {
	switch val := t.val.(type) {
	case nil:
		return 0, errors.New("不正な式")
	case nunmber:
		return int(val), nil
	case operator:
		left, err := walkTree(t.lchild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(t.rchild)
		if err != nil {
			return 0, err
		}
		return val.process(left, right), nil
	default:
		return 0, errors.New("不正な型のノード")
	}
}

func main() {
	parseTree, err := parse("5*10+20")
	if err != nil {
		panic(err)
	}
	result, err := walkTree(parseTree)
	fmt.Println(result, err)

	parseTree, err = parse("2-10*3")
	if err != nil {
		panic(err)
	}
	result, err = walkTree(parseTree)
	fmt.Println(result, err)
}

func parse(s string) (*treeNode, error) {
	if s == "5*10+20" {
		return &treeNode{
			val: operators["+"],
			lchild: &treeNode{
				val:    operators["*"],
				lchild: &treeNode{val: nunmber(5)},
				rchild: &treeNode{val: nunmber(10)},
			},
			rchild: &treeNode{val: nunmber(20)},
		}, nil
	} else if s == "2-10*3" {
		return &treeNode{
			val:    operators["-"],
			lchild: &treeNode{val: nunmber(2)},
			rchild: &treeNode{
				val:    operators["*"],
				lchild: &treeNode{val: nunmber(10)},
				rchild: &treeNode{val: nunmber(3)},
			},
		}, nil
	}
	return nil, nil
}
