package person

import "github.com/learning-go-book/circular_dependency_example/pet"

type Person struct {
	Name    string
	Age     int
	PetName string
}

var pets = map[string]pet.Pet{
	"Fluffy": {Name: "Fluffy", Type: "Cat", OwnerName: "Bob"},
	"Rex":    {Name: "Rex", Type: "Dog", OwnerName: "Julia"},
}

func (p Person) Pet() pet.Pet {
	return pets[p.PetName]
}
