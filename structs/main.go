package main

import "fmt"

type location struct {
	email   string
	zipCode int
}

type person struct {
	name     string
	lastname string
	location
}

func (p person) print() {
	fmt.Println(p.name, p.lastname)
}

func (p *person) updateName(newName string) {
	(*p).name = newName
}

/**
 * Remember, the functions make a new copy of all arguments
 * We have values by types and values by reference
 */

func main() {
	gian := person{
		name:     "Giancarlos",
		lastname: "Isasi",
		location: location{
			email:   "giancarlos.isasi@gmail.com",
			zipCode: 5154,
		},
	}

	gian.updateName("Nexus")
	gian.print()

}
