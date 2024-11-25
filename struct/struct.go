package main

import "fmt"

// defining the struct type
type Person struct{
  Name string
  Age int
}

type User struct {
	FullName string
	Age int
}

// this method becomes part of the struc type Person here p Person is called Receiver
func (p Person) logName() {
	fmt.Println(p.Name);	
}

// this method becomes part of the struc type Person here p Person is called Receiver
func (p Person) logAge() {
	fmt.Println(p.Age);	
}

// example Factory method that returns a user object
func newUser (fullName string, age int) User {
	return User{
		FullName: fullName,
		Age: age,
	}
}

// example of struct embedding, more precisely a nested struct
type Employee struct {
    Person  
    Position  string
    Salary    float64
}

// interfaces 
type Mover interface {
    Move() string
    Stop() string
}

type Car struct {
    Brand string
}

func (c Car) Move() string {
    return fmt.Sprintf("%s car is moving", c.Brand)
}

func (c Car) Stop() string {
    return fmt.Sprintf("%s car has stopped", c.Brand)
}

type Bicycle struct {
    Type string
}

func (b Bicycle) Move() string {
    return fmt.Sprintf("%s bicycle is moving", b.Type)
}

func (b Bicycle) Stop() string {
    return fmt.Sprintf("%s bicycle has stopped", b.Type)
}

func main() {

	// creating a Person type
	var person =  Person{
        Name:    "John Doe",
        Age:     30,
    }


	person.logAge()
	person.logName()

	var user User = newUser("Full Name",25);
	fmt.Println(user);

	var employee = Employee{
		Person: Person{
			Name: "My Name",
			Age: 25,
		},
		Position: "Software Engineer",
		Salary:   75000.00,
	}
	fmt.Println(employee);

	car := Car{Brand: "Toyota"}
    bike := Bicycle{Type: "Mountain"}

	fmt.Println(car.Move())
    fmt.Println(car.Stop())

	fmt.Println(bike.Move())
    fmt.Println(bike.Stop())

}