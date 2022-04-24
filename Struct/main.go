package main

import "fmt"

/*
	Interface
		- Adalah tipe data abstract, tidak memiliki implementasi langsung
		- Biasanya berisikan definisi - definisi method
		- Biasanya interface digunakan sebagai kontrak
*/

type AbstractMethod interface {
	GetFullName() string
}

type Customer struct {
	Name    *CustomerName
	Phone   string
	Address []string
}

type CustomerName struct {
	FirstName string
	LastName  string
}

// Function Method
// Make Customer as a variable then it will be function method
func (customer Customer) dialog() {
	fmt.Println("Hi, Mr/Mrs", customer.GetFullName())
}

func (customer Customer) sayName(name string) {
	fmt.Println("Hello,", name, "Nice to meet you here!")
}

func (customer Customer) GetFullName() string {
	return customer.Name.FirstName + " " + customer.Name.LastName
}

func GetFullName(abstract AbstractMethod) {
	fmt.Println("Hello ", abstract.GetFullName(), " ^_^")
}

func main() {

	customer := Customer{
		Name: &CustomerName{
			FirstName: "Joshua",
			LastName:  "Ryandafres Pangaribuan",
		},
		Phone: "082295043709",
		Address: []string{
			"Jl.Pdt Justin Sihombing, No.210",
		},
	}

	fmt.Println("Customer #1")
	fmt.Println("Customer Name\t:", customer.GetFullName())
	fmt.Println("Customer Phone\t:", customer.Phone)
	fmt.Println("Customer Address:", customer.Address[0])
	customer.dialog()
	customer.sayName("Alex Sander")
	GetFullName(customer)
}
