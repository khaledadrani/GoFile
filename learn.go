package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	name    string
	age     int
	contact contact
}

func createContactInfo() contact {
	c := contact{email: "my_email", zipCode: 1023}

	return c
}

func (p *person) updateName(newName string) {
	(*p).name = newName
}

func (p *person) updateZipCode(newZipCode int) {
	(*p).contact.zipCode = newZipCode
}

func (p person) printPerson() {
	fmt.Printf("person name is %+v \n", p)

}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func printHashMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
