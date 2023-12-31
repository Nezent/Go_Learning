package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	anon := person{firstName: "Anon", 
					lastName: "Anderson",
					contact: contactInfo{
						email: "sirajummunir31@gmail.com",
						zip: 1171,
					},
				}
	// var anon person
	// anon.firstName = "Anon"
	// anon.lastName = "Anderson"
	anonPointer := &anon
	anonPointer.updateName("Anonymous")
	anon.print()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v",p)
}
