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
	fmt.Println(anon)
	fmt.Printf("%+v",anon)
}
