package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "上杉謙信",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	var eOK Employee = m.Employee
	fmt.Println(eOK)
	// var eFail Employee := m // コンパイルエラー
	// fmt.Println(eFail)
}
