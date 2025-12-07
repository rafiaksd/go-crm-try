package main

import (
	"fmt"
	"strconv"
	"time"
	"strings"
)

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	CreatedAt time.Time
}

type CRM struct {
	customers map[int]Customer
	NextID    int
}

func NewCRM() *CRM {
	return &CRM{customers: make(map[int]Customer), NextID: 1}
}

func (crm *CRM) addCustomer(firstname, lastname,email,phone string){
	cust := Customer{
		ID:crm.NextID,
		FirstName: firstname,
		LastName: lastname,
		Email: email,
		Phone: phone,
		CreatedAt: time.Now(),
	}

	crm.customers[crm.NextID] = cust
	crm.NextID ++
	fmt.Printf("Customer {%s} - {%s} added successfully\n", strconv.Itoa(cust.ID), cust.FirstName)
}

func (crm CRM) ViewCustomer(id int){
	customer, exists := crm.customers[id]

	if !exists { fmt.Println("Customer not found!"); return}

	fmt.Println("Customer ID: ", customer.ID)
	fmt.Println("Name: ", customer.FirstName, customer.LastName)
	fmt.Println("Email:", customer.Email)
	fmt.Println("Created at:", customer.CreatedAt.Format("2006-01-02 15:04")) // 2 January 2006 3:04 pm
}

func (crm CRM) ViewAllCustomers(){
	if len(crm.customers) == 0 {
		fmt.Println("No customers foudn")
		return
	}

	for _, cust := range crm.customers {
		fmt.Printf("\nID: %d Name: %s %s Email: %s Phone: %s\n", cust.ID, cust.FirstName, cust.LastName, cust.Email, cust.Phone)
	}
} 

func (crm *CRM) UpdateCustomer(id int, firstName, lastName, email, phone string){
	customer, exists := crm.customers[id]

	if !exists { fmt.Println("Customer with ID:%d, dont exist"); return}

	customer.FirstName = firstName
	customer.LastName = lastName
	customer.Email = email
	customer.Phone = phone

	// save updated customer
	crm.customers[id]= customer

	fmt.Println("Customer - ID: %d - updated successfully", id)
}

func (crm *CRM) DeleteCustomer(id int){
	_, exists := crm.customers[id]
	if !exists { fmt.Println("Customer not found"); return }

	delete(crm.customers, id)
	fmt.Println("Customer deleted successfully")
}

func getInput(prompt string) string {
	fmt.Print(prompt)
	var input string 
	fmt.Scanln(&input)
	return input
}

func main() {
	fmt.Println("First push")
}
