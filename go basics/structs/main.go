package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func changeAge(person *Person, newAge int) {
	person.Age = newAge
}

// This is a receiver
func (p *Person) changeName(newName string) {
	p.Name = newName
}

func main() {
	fmt.Println(1)

	new_person := NewPerson("Aditya", 25)

	fmt.Printf("The new person is %+v\n", new_person)

	changeAge(&new_person, 26)
	fmt.Printf("The new age is %+v\n", new_person)

	new_person.changeName("Raksh")

	fmt.Printf("The new person is %+v\n", new_person)

}
