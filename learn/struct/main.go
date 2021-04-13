package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firtName string
	lastName string
	age      int
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstname string) {
	(*p).firtName = newFirstname
}

func main() {
	// taey := person{
	// 	firtName: "Nitipat",
	// 	lastName: "L",
	// }

	// taey := person{"Nitipat", "L"}

	// var taey person // zero value assigned
	// taey.firtName = "Nitipat"
	// taey.lastName = "L"
	// taey.age = 29
	// fmt.Println(taey)
	// fmt.Printf("%+v", taey)

	taey := person{
		firtName: "Nitipat",
		lastName: "L",
		age:      29,
		contactInfo: contactInfo{
			email:   "iamsvz@gmail.com",
			zipCode: 50200,
		},
	}
	taey.updateName("Test")
	taey.print()
}
