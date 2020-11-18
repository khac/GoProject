// create "student1" struct type with firstname, secondname, phone, email fields.
// create a variable of type slice of "student1" struct, and initialize its with 5 student records
// create "student2" struct type with fullname, phone, email fields.
// create a variable of type slice of "student2" struct,
// copy the struct variable 1 records into variable 2. also combine first name and last name to form the fullname for the variable 2)

package main

import (
	"fmt"
)

type student1 struct {
	firstname  string
	secondname string
	phone      string
	email      string
}

type student2 struct {
	fullname string
	phone    string
	email    string
}

func main() {
	var m [5]student1 = [5]student1{}
	m[0] = student1{firstname: "a", secondname: "f", phone: "123-456", email: "a@f.com"}
	m[1] = student1{firstname: "b", secondname: "g", phone: "987-126", email: "b@g.com"}
	m[2] = student1{firstname: "c", secondname: "h", phone: "873-945", email: "c@h.com"}
	m[3] = student1{firstname: "d", secondname: "i", phone: "983-356", email: "d@i.com"}
	m[4] = student1{firstname: "e", secondname: "j", phone: "777-443", email: "e@j.com"}

	n := [5]student2{}
	for i, _ := range m {
		n[i] = student2{fullname: m[i].firstname + " " + m[i].secondname, phone: m[i].phone, email: m[i].email}
		fmt.Println(n[i])
	}

}
